package broker

import (
	"fmt"

	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/iggy/adapters/broker"
	"github.com/kaibling/iggy/adapters/broker/loopback"
	"github.com/kaibling/iggy/adapters/broker/nats"
	"github.com/kaibling/iggy/apperror"
)

func NewSubscriber(subConfig broker.SubscriberConfig, //nolint:ireturn,nolintlint
	logger logging.Writer,
) (broker.Subscriber, error) {
	l := logger.NewScope("subscriber")

	switch subConfig.Config.Broker.BrokerName {
	case "rabbitMQ":
		panic("not implemented")
	case "nats":
		return nats.NewNATSClient(subConfig, l)
	case "loopback":
		return loopback.NewLoopback(subConfig, l), nil
	default:
		return nil, apperror.NewStringGeneric(fmt.Sprintf("subscriber adapter %s not found", subConfig.Config.Broker.BrokerName))
	}
}

func NewPublisher(subConfig broker.SubscriberConfig, logger logging.Writer, //nolint:ireturn,nolintlint
) (broker.Publisher, error) {
	l := logger.NewScope("publisher")

	switch subConfig.Config.Broker.BrokerName {
	case "rabbitMQ":
		panic("not implemented")
	case "nats":
		return nats.NewNATSClient(subConfig, l)
	case "loopback":
		return loopback.NewLoopback(subConfig, l), nil
	default:
		return nil, apperror.NewStringGeneric(fmt.Sprintf("publisher adapter %s not found", subConfig.Config.Broker.BrokerName))
	}
}
