package pkg

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	"go.uber.org/zap"
	"golang.org/x/sys/unix"
	"golang.zx2c4.com/wireguard/wgctrl"
	"kernel.org/pub/linux/libs/security/libcap/cap"
	"riasc.eu/wice/pkg/config"
	"riasc.eu/wice/pkg/core"
	"riasc.eu/wice/pkg/device"
	errs "riasc.eu/wice/pkg/errors"
	ac "riasc.eu/wice/pkg/feat/auto"
	ep "riasc.eu/wice/pkg/feat/disc/ep"
	cs "riasc.eu/wice/pkg/feat/sync/config"
	rs "riasc.eu/wice/pkg/feat/sync/routes"
	"riasc.eu/wice/pkg/util"
	"riasc.eu/wice/pkg/watcher"
	"riasc.eu/wice/pkg/wg"

	"riasc.eu/wice/pkg/signaling"

	"go.uber.org/zap/zapio"
)

type Daemon struct {
	*watcher.Watcher

	// Features

	AutoConfig        *ac.AutoConfiguration
	ConfigSync        *cs.ConfigSynchronization
	RouteSync         *rs.RouteSynchronization
	EndpointDiscovery *ep.EndpointDiscovery

	// Shared

	Backend *signaling.MultiBackend
	client  *wgctrl.Client
	config  *config.Config

	stop    chan any
	signals chan os.Signal

	logger *zap.Logger
}

func NewDaemon(cfg *config.Config) (*Daemon, error) {
	var err error

	logger := zap.L().Named("daemon")

	// Check permissions
	if !util.HasCapabilities(cap.NET_ADMIN) {
		return nil, errors.New("insufficient privileges. Pleas run wice as root user or with NET_ADMIN capabilities")
	}

	// Create backend

	urls := []*url.URL{}
	for _, u := range cfg.Backends {
		urls = append(urls, &u.URL)
	}

	backend, err := signaling.NewMultiBackend(urls, &signaling.BackendConfig{
		OnReady: []signaling.BackendReadyHandler{},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to initialize signaling backend: %w", err)
	}

	// Create Wireguard netlink socket
	client, err := wgctrl.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create Wireguard client: %w", err)
	}

	d := &Daemon{
		config:  cfg,
		client:  client,
		Backend: backend,

		stop:    make(chan any),
		signals: SetupSignals(),

		logger: logger,
	}

	if d.Watcher, err = watcher.New(d.client, cfg.WatchInterval, &cfg.Wireguard.InterfaceFilter.Regexp); err != nil {
		return nil, fmt.Errorf("failed to initialize watcher: %w", err)
	}

	// Check if Wireguard interface can be created by the kernel
	if !cfg.Wireguard.Userspace {
		cfg.Wireguard.Userspace = !wg.KernelModuleExists()
	}

	if err := d.setupFeatures(); err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Daemon) setupFeatures() error {
	var err error

	if d.config.AutoConfig.Enabled {
		if d.AutoConfig, err = ac.New(d.Watcher, d.client); err != nil {
			return fmt.Errorf("failed to start interface auto configuration: %w", err)
		}
	}

	if d.config.ConfigSync.Enabled {
		if d.ConfigSync, err = cs.New(d.Watcher, d.client,
			d.config.ConfigSync.Path,
			d.config.ConfigSync.Watch,
			d.config.Wireguard.Userspace); err != nil {

			return fmt.Errorf("failed to start configuration file synchronization: %w", err)
		}

		d.logger.Info("Started configuration file synchronization")
	}

	if d.config.RouteSync.Enabled {
		if d.RouteSync, err = rs.New(d.Watcher, d.config.RouteSync.Table); err != nil {
			return fmt.Errorf("failed to start allowed-ips <-> kernel route synchronization: %w", err)
		}

		d.logger.Info("Started allowed-ips <-> kernel route synchronization")
	}

	if d.config.EndpointDisc.Enabled {
		if d.EndpointDiscovery, err = ep.New(d.Watcher, d.config, d.client, d.Backend); err != nil {
			return fmt.Errorf("failed to start endpoint discovery: %w", err)
		}

		d.logger.Info("Started endpoint discovery")
	}

	return nil
}

func (d *Daemon) Run() {
	if err := d.CreateInterfacesFromArgs(); err != nil {
		d.logger.Fatal("failed to create interfaces", zap.Error(err))
	}

	go d.Watcher.Run()

out:
	for {
		select {
		case sig := <-d.signals:
			d.logger.Debug("Received signal", zap.String("signal", sig.String()))
			switch sig {
			case unix.SIGUSR1:
				if err := d.Sync(); err != nil {
					d.logger.Error("Failed to synchronize interfaces", zap.Error(err))
				}
			default:
				break out
			}

		case <-d.stop:
			break out
		}
	}
}

func (d *Daemon) IsRunning() bool {
	select {
	case _, running := <-d.stop:
		return running
	default:
		return true
	}
}

func (d *Daemon) Stop() error {
	if !d.IsRunning() {
		return errs.ErrAlreadyStopped
	}

	close(d.stop)

	return nil
}

func (d *Daemon) Close() error {
	if err := d.Stop(); err != nil && !errors.Is(err, errs.ErrAlreadyStopped) {
		return err
	}

	if err := d.Watcher.Close(); err != nil {
		return fmt.Errorf("failed to close interface: %w", err)
	}

	if err := d.client.Close(); err != nil {
		return fmt.Errorf("failed to close Wireguard client: %w", err)
	}

	return nil
}

func (d *Daemon) CreateInterfacesFromArgs() error {
	var devs device.Devices
	devs, err := d.client.Devices()
	if err != nil {
		return err
	}

	for _, intfName := range d.config.Wireguard.Interfaces {
		dev := devs.GetByName(intfName)
		if dev != nil {
			d.logger.Warn("Interface already exists. Skipping..", zap.Any("intf", intfName))
			continue
		}

		i, err := core.CreateInterface(intfName, d.config.Wireguard.Userspace, d.client)
		if err != nil {
			return fmt.Errorf("failed to create Wireguard device: %w", err)
		}

		if d.logger.Core().Enabled(zap.DebugLevel) {
			d.logger.Debug("Initialized interface:")
			if err := i.DumpConfig(&zapio.Writer{Log: d.logger}); err != nil {
				return err
			}
		}

		d.Watcher.Interfaces[i.Name()] = i
	}

	return nil
}
