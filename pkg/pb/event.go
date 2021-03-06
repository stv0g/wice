package pb

import (
	"strings"

	"go.uber.org/zap"
	"riasc.eu/wice/pkg/crypto"
)

func (e *Event) Log(l *zap.Logger, msg string, fields ...zap.Field) {
	if e.Type != Event_UNKNOWN {
		fields = append(fields, zap.String("type", strings.ToLower(e.Type.String())))
	}

	if e.Event != nil {
		fields = append(fields, zap.Any("event", e.Event))
	}

	if e.Interface != "" {
		fields = append(fields, zap.String("interface", e.Interface))
	}

	if e.Peer != nil {
		pk, _ := crypto.ParseKeyBytes(e.Peer)
		fields = append(fields, zap.Any("peer", pk))
	}

	if e.Time != nil {
		fields = append(fields, zap.Time("time", e.Time.Time()))
	}

	l.Info(msg, fields...)
}
