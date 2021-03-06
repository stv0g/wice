package inprocess

import (
	"context"

	"go.uber.org/zap"
	"riasc.eu/wice/pkg/crypto"
	"riasc.eu/wice/pkg/pb"
	"riasc.eu/wice/pkg/signaling"
)

var (
	subs = signaling.NewSubscriptionsRegistry()
)

func init() {
	signaling.Backends["inprocess"] = &signaling.BackendPlugin{
		New:         NewBackend,
		Description: "In-Process",
	}
}

type Backend struct {
	logger *zap.Logger
}

func NewBackend(cfg *signaling.BackendConfig, logger *zap.Logger) (signaling.Backend, error) {
	b := &Backend{
		logger: logger,
	}

	for _, h := range cfg.OnReady {
		h.OnSignalingBackendReady(b)
	}

	return b, nil
}

func (b *Backend) Type() pb.BackendType {
	return pb.BackendType_INPROCESS
}

func (b *Backend) SubscribeAll(ctx context.Context, kp *crypto.Key, h signaling.MessageHandler) error {
	_, err := subs.SubscribeAll(kp, h)
	return err
}

func (b *Backend) Subscribe(ctx context.Context, kp *crypto.KeyPair, h signaling.MessageHandler) error {
	_, err := subs.Subscribe(kp, h)
	return err
}

func (b *Backend) Publish(ctx context.Context, kp *crypto.KeyPair, msg *signaling.Message) error {
	env, err := msg.Encrypt(kp)
	if err != nil {
		return err
	}

	return subs.NewMessage(env)
}

func (b *Backend) Close() error {
	return nil
}
