package rabbit

import (
	"context"
	"sync"
	"time"

	"github.com/bytedance/sonic"
	"github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	uri          string
	connection   *amqp091.Connection
	ChannelPools sync.Pool
}

func NewRabbit(rabbitmqUrl string) *Rabbit {
	return &Rabbit{
		uri: rabbitmqUrl,
	}
}

func (r *Rabbit) Init() *Rabbit {
	r.connection = getConnection(r.uri)
	r.ChannelPools = sync.Pool{
		New: func() any {
			ch, err := r.connection.Channel()
			if err != nil {
				panic(err)
			}

			return ch
		},
	}

	r.setupExchanges()

	return r
}

func (r *Rabbit) setupExchanges() {
	ch := r.ChannelPools.Get().(*amqp091.Channel)

	ch.ExchangeDeclare(EXCHANGE_USER_SERVICE, "direct", true, true, false, false, nil)
}

func (r *Rabbit) SendToExchange(routingKey string, data any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	json, err := sonic.Marshal(data)
	if err != nil {
		return err
	}

	ch := r.ChannelPools.Get().(*amqp091.Channel)
	if err := ch.PublishWithContext(ctx, EXCHANGE_USER_SERVICE, routingKey, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        json,
	}); err != nil {
		return err
	}

	return nil
}

func (r *Rabbit) Close() error {
	if err := r.connection.Close(); err != nil {
		return err
	}

	return nil
}
