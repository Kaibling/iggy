package nats

import (
	"fmt"

	"github.com/kaibling/apiforge/logging"
	"github.com/nats-io/nats-server/v2/server"
)

func Run(log logging.Writer) {
	opts := &server.Options{
		Host: "0.0.0.0",
		Port: 4222,
	}

	ns, err := server.NewServer(opts)
	if err != nil {
		log.ErrorMsg(fmt.Sprintf("Error starting NATS server: %v", err))
	}

	log.Info(fmt.Sprintf("Embedded NATS server started on %s:%d", opts.Host, opts.Port))
	ns.Start()
}
