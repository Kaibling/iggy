package broker

import (
	"fmt"

	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/iggy/adapters/broker"
	"github.com/kaibling/iggy/adapters/broker/loopback"
	"github.com/kaibling/iggy/adapters/broker/nats"
	"github.com/kaibling/iggy/apperror"
	"github.com/kaibling/iggy/service"
)

func NewSubscriber(cfg service.Config, //nolint:ireturn,nolintlint
	logger logging.Writer,
) (broker.Subscriber, error) {
	l := logger.NewScope("subscriber")

	switch cfg.Config.Broker.BrokerName {
	case "rabbitMQ":
		panic("not implemented")
	case "nats":
		return nats.NewNATSClient(cfg, l)
	case "loopback":
		return loopback.NewLoopback(cfg, l), nil
	default:
		return nil, apperror.NewStringGeneric(fmt.Sprintf("subscriber adapter %s not found", cfg.Config.Broker.BrokerName))
	}
}

func NewPublisher(cfg service.Config, logger logging.Writer, //nolint:ireturn,nolintlint
) (broker.Publisher, error) {
	l := logger.NewScope("publisher")

	switch cfg.Config.Broker.BrokerName {
	case "rabbitMQ":
		panic("not implemented")
	case "nats":
		return nats.NewNATSClient(cfg, l)
	case "loopback":
		return loopback.NewLoopback(cfg, l), nil
	default:
		return nil, apperror.NewStringGeneric(fmt.Sprintf("publisher adapter %s not found", cfg.Config.Broker.BrokerName))
	}
}
