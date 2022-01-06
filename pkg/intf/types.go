package intf

import "strings"

type PeerModifier int

const (
	PeerModifiedNone              PeerModifier = 0
	PeerModifiedEndpoint          PeerModifier = (1 << 0)
	PeerModifiedKeepaliveInterval PeerModifier = (1 << 1)
	PeerModifiedProtocolVersion   PeerModifier = (1 << 2)
	PeerModifiedAllowedIPs        PeerModifier = (1 << 3)
	PeerModifiedHandshakeTime     PeerModifier = (1 << 4)
)

var (
	peerModifiersStrings = []string{
		"endpoint",
		"keepalive-interval",
		"protocol-version",
		"allowed-ips",
		"handshake-time",
	}
)

func (i PeerModifier) String() string {
	modifiers := []string{}

	for j := 0; j <= 4; i++ {
		if i&(1<<j) != 0 {
			modifiers = append(modifiers, peerModifiersStrings[j])
		}
	}

	return strings.Join(modifiers, ",")
}