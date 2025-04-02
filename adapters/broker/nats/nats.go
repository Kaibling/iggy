package nats

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/iggy/bootstrap"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/utility"
	"github.com/kaibling/iggy/service"
	nats_go "github.com/nats-io/nats.go"
)

const waitingDuration = 100 * time.Second
const logScope = "subscriber"

func NewNATSClient(cfg service.Config, l logging.Writer) (*Client, error) {
	nc, err := nats_go.Connect(cfg.Config.Broker.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to NATS server: %w", err)
	}

	return &Client{cfg: cfg, l: l.NewScope(logScope), nc: nc, backupLogger: l.NewScope(logScope)}, nil
}

type Client struct {
	cfg          service.Config
	l            logging.Writer
	backupLogger logging.Writer
	nc           *nats_go.Conn
}

func (n *Client) Publish(_ context.Context, channelName string, message []byte) error {
	err := n.nc.Publish(channelName, message)
	if err != nil {
		return fmt.Errorf("error publishing message: %w", err)
	}

	n.l.Info("Message published to" + channelName)

	return nil
}

func (n *Client) Subscribe(ctx context.Context, channelName string) error {
	// Subscribe synchronously
	sub, err := n.nc.SubscribeSync(channelName)
	if err != nil {
		return fmt.Errorf("error subscribing to subject: %w", err)
	}

	n.l.Info("Subscribed to subject:" + channelName)

	// Process messages
	for {
		cycleID := utils.NewULID().String()
		n.l = n.backupLogger.NewScope(logScope)
		n.l.AddStringField("cycle_id", cycleID)
		msg, err := sub.NextMsg(waitingDuration)
		if err != nil {
			if errors.Is(err, nats_go.ErrTimeout) {
				n.l.Debug("listening timeout: no messages received")
			} else {
				n.l.Info(fmt.Sprintf("mesasge read error occurred: %v", err))
			}

			continue
		}

		n.l.Debug("received message")

		t, err := utility.DecodeToStruct[entity.Task](msg.Data)
		if err != nil {
			n.l.Error(err)

			continue
		}
		n.l.Debug(fmt.Sprintf("received Workflow %s", t.WorkflowID))

		n.l.AddAnyField("request_id", t.RequestID)
		n.cfg.Log = n.l

		if err := bootstrap.WorkerExecution(ctx, n.cfg, t); err != nil {
			if strings.Contains(err.Error(), "not found") {
				n.l.Warn(err.Error())
			} else {
				n.l.Error(err)
			}

		}
	}
}
