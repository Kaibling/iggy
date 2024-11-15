package loopback

import (
	"context"

	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/iggy/adapters/broker"
	"github.com/kaibling/iggy/bootstrap"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/utility"
)

var LoopbackChannel = make(chan []byte) //nolint: gochecknoglobals

func NewLoopback(subConfig broker.SubscriberConfig, l logging.Writer) *Loopback {
	return &Loopback{subConfig, l.NewScope("Subscriber")}
}

type Loopback struct {
	cfg broker.SubscriberConfig
	l   logging.Writer
}

func (l *Loopback) Subscribe(ctx context.Context, _ string) error {
	// TODO multi worker goroutine
	l.l.Info("loopback worker waiting...")

	for {
		select {
		case newMessage := <-LoopbackChannel:
			t, err := utility.DecodeToStruct[entity.Task](newMessage)
			if err != nil {
				l.l.Error(err)
			}
			l.l.AddAnyField("request_id", t.RequestID)

			if err := bootstrap.WorkerExecution(ctx, l.cfg, t); err != nil {
				l.l.Error(err)
			}
		case <-ctx.Done():
			l.l.Info("shuting down loopback worker")

			return ctx.Err()
		}
	}
}

func (l *Loopback) Publish(_ context.Context, _ string, message []byte) error {
	LoopbackChannel <- message

	return nil
}
