package nats

import (
	"context"
	"fmt"
	"time"

	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/iggy/adapters/broker"
	"github.com/kaibling/iggy/bootstrap"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/utility"
	nats_go "github.com/nats-io/nats.go"
)

func NewNATSClient(cfg broker.SubscriberConfig, l logging.Writer) (*NATSClient, error) {
	url := "nats://0.0.0.0:4222" // TODO
	nc, err := nats_go.Connect(url)
	if err != nil {
		return nil, fmt.Errorf("error connecting to NATS server: %v", err)
	}
	return &NATSClient{cfg: cfg, l: l.NewScope("Subscriber"), nc: nc}, nil
}

type NATSClient struct {
	cfg broker.SubscriberConfig
	l   logging.Writer
	nc  *nats_go.Conn
}

func (n *NATSClient) Publish(ctx context.Context, channelName string, message []byte) error {
	err := n.nc.Publish(channelName, []byte(message))
	if err != nil {
		return fmt.Errorf("error publishing message: %v", err)
	}
	n.l.Info(fmt.Sprintf("Message published to %s", channelName))
	return nil
}

func (n *NATSClient) Subscribe(ctx context.Context, channelName string) error {
	// Subscribe synchronously

	sub, err := n.nc.SubscribeSync(channelName)
	if err != nil {
		return fmt.Errorf("error subscribing to subject: %v", err)
	}

	n.l.Info(fmt.Sprintf("Subscribed to subject: %s", channelName))

	//Process messages
	for {
		msg, err := sub.NextMsg(10 * time.Second) // Wait up to 10 seconds for a message
		if err != nil {
			n.l.Info(fmt.Sprintf("no messages received or error occurred: %v", err))
			continue
		}
		n.l.Info("Received message")
		t, err := utility.DecodeToStruct[entity.Task](msg.Data)
		if err != nil {
			n.l.Error(err)
		}

		n.l.AddAnyField("request_id", t.RequestID)
		n.cfg.Log = n.l

		if err := bootstrap.WorkerExecution(ctx, n.cfg, t); err != nil {
			n.l.Error(err)
		}
	}

	// for {
	// 	select {
	// 	case newMessage := <-LoopbackChannel:
	// 		t, err := utility.DecodeToStruct[entity.Task](newMessage)
	// 		if err != nil {
	// 			l.l.Error(err)
	// 		}

	// 		l.l.AddAnyField("request_id", t.RequestID)
	// 		l.cfg.Log = l.l

	// 		if err := bootstrap.WorkerExecution(ctx, l.cfg, t); err != nil {
	// 			l.l.Error(err)
	// 		}

	// 	case <-ctx.Done():
	// 		l.l.Info("shuting down loopback worker")

	// 		return ctx.Err()
	// 	}
}
