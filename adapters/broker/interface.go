package broker

import (
	"context"
)

type Publisher interface {
	Publish(ctx context.Context, channelName string, message []byte) error
}

type Subscriber interface {
	Subscribe(ctx context.Context, channelName string) error
}

// type SubscriberConfig struct {
// 	Config   config.Configuration
// 	Username string
// 	DBPool   *pgxpool.Pool
// 	Log      logging.Writer
// }
