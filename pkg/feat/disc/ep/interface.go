package ep

import (
	"fmt"

	"github.com/pion/ice/v2"
	"go.uber.org/zap"
	"riasc.eu/wice/pkg/core"
	"riasc.eu/wice/pkg/errors"
	"riasc.eu/wice/pkg/proxy"
)

type Interface struct {
	*core.Interface

	Discovery *EndpointDiscovery

	nat *proxy.NAT

	udpMux      ice.UDPMux
	udpMuxSrflx ice.UniversalUDPMux

	logger *zap.Logger
}

func NewInterface(ci *core.Interface, d *EndpointDiscovery) (*Interface, error) {
	i := &Interface{
		Interface: ci,
		Discovery: d,

		logger: zap.L().Named("ice.interface"),
	}

	// Create per-interface UDPMux
	var err error
	var lPortHost, lPortSrfx int

	if i.udpMux, lPortHost, err = proxy.CreateUDPMux(i.ListenPort); err != nil && err != errors.ErrNotSupported {
		return nil, fmt.Errorf("failed to setup host UDP mux: %w", err)
	}

	if i.udpMuxSrflx, lPortSrfx, err = proxy.CreateUDPMuxSrflx(); err != nil && err != errors.ErrNotSupported {
		return nil, fmt.Errorf("failed to setup srflx UDP mux: %w", err)
	}

	// Setup NAT
	ident := fmt.Sprintf("wice-if%d", i.KernelDevice.Index())
	if i.nat, err = proxy.NewNAT(ident); err != nil && err != errors.ErrNotSupported {
		return nil, fmt.Errorf("failed to setup NAT: %w", err)
	}

	// Redirect non-STUN traffic directed at UDP muxes to Wireguard interface via in-kernel port redirect / NAT
	if err := i.nat.RedirectNonSTUN(lPortHost, i.ListenPort); err != nil {
		return nil, fmt.Errorf("failed to setup port redirect for server reflexive UDP mux: %w", err)
	}

	if err := i.nat.RedirectNonSTUN(lPortSrfx, i.ListenPort); err != nil {
		return nil, fmt.Errorf("failed to setup port redirect for server reflexive UDP mux: %w", err)
	}

	return i, nil
}

func (i *Interface) Close() error {
	if err := i.nat.Close(); err != nil {
		return fmt.Errorf("failed to de-initialize NAT: %w", err)
	}

	if err := i.udpMux.Close(); err != nil {
		return fmt.Errorf("failed to do-initialize UDP mux: %w", err)
	}

	if err := i.udpMuxSrflx.Close(); err != nil {
		return fmt.Errorf("failed to do-initialize srflx UDP mux: %w", err)
	}

	return nil
}
