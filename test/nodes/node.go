//go:build linux

package nodes

import (
	g "github.com/stv0g/gont/pkg"
)

type Node interface {
	g.Node

	Start(args ...any) error
	Stop() error
	Close() error
}
