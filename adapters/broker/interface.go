package broker

type BrokerAdapter interface {
	Subscribe(channelName string) error
	Publish(channelName string, message []byte) error
}
