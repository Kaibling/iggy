package broker

import (
	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/iggy/pkg/broker/nats"
)

func StartServer(log logging.Writer) {
	go nats.Run(log.NewScope("broker"))
}
