package broker

import (
	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/iggy/pkg/broker/nats"
	"github.com/kaibling/iggy/pkg/config"
)

func StartServer(log logging.Writer, cfg config.Configuration) {
	go nats.Run(log.NewScope("broker"), cfg)
}
