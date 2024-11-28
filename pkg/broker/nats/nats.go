package nats

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/iggy/pkg/config"
	"github.com/nats-io/nats-server/v2/server"
)

func Run(log logging.Writer, cfg config.Configuration) {
	opts := parseConnString(cfg.Broker.ConnectionString)

	ns, err := server.NewServer(opts)
	if err != nil {
		log.ErrorMsg(fmt.Sprintf("Error starting NATS server: %v", err))
	}

	log.Info(fmt.Sprintf("Embedded NATS server started on %s:%d", opts.Host, opts.Port))
	ns.Start()
}

func parseConnString(connStr string) *server.Options {
	hostPort := strings.SplitAfter(connStr, "//")
	split := strings.Split(hostPort[1], ":")
	host := split[0]
	port, _ := strconv.Atoi(split[1])

	return &server.Options{
		Host: host,
		Port: port,
	}
}
