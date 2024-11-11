package broker

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaibling/iggy/pkg/config"
)

type Publisher interface {
	Publish(ctx context.Context, channelName string, message []byte) error
}

type Subscriber interface {
	Subscribe(ctx context.Context, channelName string) error
}

type SubscriberConfig struct {
	Config   config.Configuration
	Username string
	DBPool   *pgxpool.Pool
}
