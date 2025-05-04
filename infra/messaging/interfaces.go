package messaging

type MessagePublisherInterface interface {
	Publish(exchange, routingKey string, body []byte) error
	Consumer(handler func(string), exchange string) error
}
