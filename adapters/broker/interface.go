package broker

type Adapter interface {
	Subscribe(channelName string) error
	Publish(channelName string, message []byte) error
}
