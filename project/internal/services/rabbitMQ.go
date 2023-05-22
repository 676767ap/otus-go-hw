package services

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/676767ap/project/internal/config"
	"github.com/676767ap/project/internal/entity"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	EventTypeImpression = iota
	EventTypeClick
)

type Exchange interface {
	DeclareQueue(ctx context.Context, name string) (Queue amqp.Queue, err error)
	Consume(ctx context.Context, queueName, consumerName string) (<-chan amqp.Delivery, error)
	Close(ctx context.Context) error
	SendEvent(ctx context.Context, queue string, event *entity.Stat) error
}

var errTimeoutClosingRabbitMQ = errors.New("timeout closing rabbitmq")

type ExchangeSettings struct {
	Kind                                  string
	Durable, AutoDelete, Internal, NoWait bool
}

type QueueSettings struct {
	Durable, AutoDelete, Exclusive, NoWait, BindNoWait bool
	BindingKey                                         string
}

type PublisherSettings struct {
	Mandatory, Immediate bool
	RoutingKey           string
}

type ConsumerSettings struct {
	AutoAck, Exclusive, NoLocal, NoWait bool
}

type Queue struct {
	Queue             amqp.Queue
	Conn              *amqp.Connection
	Ch                *amqp.Channel
	ExchangeSettings  ExchangeSettings
	QueueSettings     QueueSettings
	PublisherSettings PublisherSettings
	ConsumerSettings  ConsumerSettings
}

func New(conn *amqp.Connection, ch *amqp.Channel) Queue {
	return Queue{
		Conn: conn,
		Ch:   ch,
	}
}

func (q Queue) DeclareExchange(_ context.Context, name string) error {
	return q.Ch.ExchangeDeclare(
		name,
		q.ExchangeSettings.Kind,
		q.ExchangeSettings.Durable,
		q.ExchangeSettings.AutoDelete,
		q.ExchangeSettings.Internal,
		q.ExchangeSettings.NoWait,
		nil,
	)
}

func (q Queue) DeclareQueue(_ context.Context, name string) (Queue amqp.Queue, err error) {
	q.Queue, _ = q.Ch.QueueDeclare(
		name,
		q.QueueSettings.Durable,
		q.QueueSettings.AutoDelete,
		q.QueueSettings.Exclusive,
		q.QueueSettings.NoWait,
		nil,
	)
	return
}

func (q Queue) BindQueue(_ context.Context, queueName, exchangeName string) error {
	return q.Ch.QueueBind(
		queueName,
		q.QueueSettings.BindingKey,
		exchangeName,
		q.QueueSettings.BindNoWait,
		nil,
	)
}

func (q Queue) Consume(_ context.Context, queueName, consumerName string) (<-chan amqp.Delivery, error) {
	return q.Ch.Consume(
		queueName,
		consumerName,
		q.ConsumerSettings.AutoAck,
		q.ConsumerSettings.Exclusive,
		q.ConsumerSettings.NoLocal,
		q.ConsumerSettings.NoWait,
		nil,
	)
}

func (q Queue) SendEvent(ctx context.Context, queue string, event *entity.Stat) (err error) {
	var bytes []byte
	if bytes, err = json.Marshal(event); err != nil {
		return err
	}
	return q.Ch.PublishWithContext(
		ctx,
		queue,
		q.PublisherSettings.RoutingKey,
		q.PublisherSettings.Mandatory,
		q.PublisherSettings.Immediate,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bytes,
		})
}

func (q Queue) Close(ctx context.Context) error {
	var (
		cnt            = 0
		ch             = make(chan error, 2)
		newCtx, cancel = context.WithTimeout(ctx, 3*time.Second)
	)
	defer cancel()

	go func() {
		ch <- q.Ch.Close()
		ch <- q.Conn.Close()
	}()

	for {
		select {
		case <-newCtx.Done():
			return errTimeoutClosingRabbitMQ
		case err := <-ch:
			if err != nil {
				return err
			}
			cnt++
			if cnt == 2 {
				return err
			}
		}
	}
}

func NewQueue(ctx context.Context, cfg *config.Config) (Exchange, error) {
	var queue Exchange

	conn, err := amqp.Dial(cfg.BuildDSNRabbitMQ())
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	rabbit := New(conn, ch)
	rabbit.QueueSettings = QueueSettings{
		Durable:    cfg.RabbitMQ.Queue.Durable,
		AutoDelete: cfg.RabbitMQ.Queue.AutoDelete,
		Exclusive:  cfg.RabbitMQ.Queue.Exclusive,
		NoWait:     cfg.RabbitMQ.Queue.NoWait,
		BindNoWait: cfg.RabbitMQ.Queue.BindNoWait,
		BindingKey: cfg.RabbitMQ.Queue.BindingKey,
	}
	rabbit.PublisherSettings = PublisherSettings{
		Mandatory:  cfg.RabbitMQ.Publisher.Mandatory,
		Immediate:  cfg.RabbitMQ.Publisher.Immediate,
		RoutingKey: cfg.RabbitMQ.Publisher.RoutingKey,
	}
	rabbit.ExchangeSettings = ExchangeSettings{
		Kind:       cfg.RabbitMQ.Exchange.Kind,
		Durable:    cfg.RabbitMQ.Exchange.Durable,
		AutoDelete: cfg.RabbitMQ.Exchange.AutoDelete,
		Internal:   cfg.RabbitMQ.Exchange.Internal,
		NoWait:     cfg.RabbitMQ.Exchange.NoWait,
	}
	rabbit.ConsumerSettings = ConsumerSettings{
		AutoAck:   cfg.RabbitMQ.Consumer.AutoAck,
		Exclusive: cfg.RabbitMQ.Consumer.Exclusive,
		NoLocal:   cfg.RabbitMQ.Consumer.NoLocal,
		NoWait:    cfg.RabbitMQ.Consumer.NoWait,
	}
	if err := rabbit.DeclareExchange(ctx, cfg.RabbitMQ.Exchange.Name); err != nil {
		return nil, err
	}
	queue = rabbit

	return queue, nil
}
