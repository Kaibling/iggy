package broker

import (
	"fmt"

	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/iggy/adapters/broker"
	"github.com/kaibling/iggy/adapters/broker/loopback"
	"github.com/kaibling/iggy/apperror"
)

func NewSubscriber(subConfig broker.SubscriberConfig, //nolint:ireturn,nolintlint
	workerName string, logger logging.Writer,
) (broker.Subscriber, error) {
	l := logger.NewScope("subscriber")

	switch workerName {
	case "rabbitMQ":
		panic("not implemented")
	case "loopback":
		return loopback.NewLoopback(subConfig, l), nil
	default:
		return nil, apperror.NewStringGeneric(fmt.Sprintf("subscriber adapter %s not found", workerName))
	}
}

func NewPublisher(subConfig broker.SubscriberConfig, //nolint:ireturn,nolintlint
	workerName string, logger logging.Writer,
) (broker.Publisher, error) {
	l := logger.NewScope("publisher")

	switch workerName {
	case "rabbitMQ":
		panic("not implemented")
	case "loopback":
		return loopback.NewLoopback(subConfig, l), nil
	default:
		return nil, apperror.NewStringGeneric(fmt.Sprintf("publisher adapter %s not found", workerName))
	}
}
