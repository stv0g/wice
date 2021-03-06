package watcher

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"
)

func normalizeSocketName(name string) string {
	name = path.Base(name)
	return strings.TrimSuffix(name, ".sock")
}

func (w *Watcher) watchUser() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to create fsnotify watcher: %w", err)
	}

	if _, err := os.Stat(wireguardSockDir); !os.IsNotExist(err) {
		if err := watcher.Add(wireguardSockDir); err != nil {
			return fmt.Errorf("failed to watch %s: %w", wireguardSockDir, err)
		}
	}

	go func() {
		w.logger.Debug("Start watching for changes of Wireguard userspace devices")

	out:
		for {
			select {

			// Fsnotify events
			case event := <-watcher.Events:
				w.logger.Debug("Received fsnotify event", zap.Any("event", event))

				name := normalizeSocketName(event.Name)

				if event.Op&fsnotify.Create == fsnotify.Create {
					w.events <- InterfaceEvent{
						Op:   InterfaceAdded,
						Name: name,
					}
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					w.events <- InterfaceEvent{
						Op:   InterfaceDeleted,
						Name: name,
					}
				} else {
					w.logger.Warn("Unknown fsnotify event", zap.Any("event", event))
				}

			// Fsnotify errors
			case w.errors <- <-watcher.Errors:
				w.logger.Debug("Error while watching for link changes")

			case <-w.stop:
				break out
			}
		}

		w.logger.Debug("Stop watching for changes of Wireguard userspace devices")
	}()

	return nil
}
