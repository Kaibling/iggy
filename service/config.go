package service

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/iggy/pkg/config"
)

type Config struct {
	Config   config.Configuration
	Username string
	DBPool   *pgxpool.Pool
	Log      logging.Writer
}
