package rabbit

import "github.com/rabbitmq/amqp091-go"

func getConnection(url string) *amqp091.Connection {
	conn, err := amqp091.Dial(url)

	if err != nil {
		panic(err)
	}

	return conn
}
