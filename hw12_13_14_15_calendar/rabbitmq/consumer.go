package rabbitmq

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/cenkalti/backoff"
	"github.com/streadway/amqp"
)

type AMQPConsumer interface {
	Handle(fn func(<-chan amqp.Delivery)) error
}

type Consumer struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	done         chan error
	consumerTag  string
	uri          string
	exchangeName string
	exchangeType string
	queue        string
	bindingKey   string
	maxInterval  time.Duration
}

func NewConsumer(consumerTag, uri, exchangeName, exchangeType, queue, bindingKey string, maxInt time.Duration) *Consumer {
	return &Consumer{
		consumerTag:  consumerTag,
		uri:          uri,
		exchangeName: exchangeName,
		exchangeType: exchangeType,
		queue:        queue,
		bindingKey:   bindingKey,
		done:         make(chan error),
		maxInterval:  maxInt,
	}
}

func (c *Consumer) connect() error {
	var err error

	c.conn, err = amqp.Dial(c.uri)
	handleError(err, "Can't connect to AMQP")

	c.channel, err = c.conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	go func() {
		logrus.Printf("closing: %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
		c.done <- errors.New("channel closed")
	}()

	if err = c.channel.ExchangeDeclare(
		c.exchangeName,
		c.exchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return errors.Errorf("exchange declare: %s", err)
	}

	return nil
}

func (c *Consumer) announceQueue() (<-chan amqp.Delivery, error) {
	queue, err := c.channel.QueueDeclare(
		c.queue,
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, fmt.Sprintf("Could not declare `%s` queue", c.queue))

	err = c.channel.Qos(50, 0, false)
	handleError(err, "Could not configure QoS")

	if err = c.channel.QueueBind(
		queue.Name,
		c.bindingKey,
		c.exchangeName,
		false,
		nil,
	); err != nil {
		return nil, errors.Errorf("queue dind: %s", err)
	}

	msgs, err := c.channel.Consume(
		queue.Name,
		c.consumerTag,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, errors.Errorf("queue consume: %s", err)
	}

	return msgs, nil
}

func (c *Consumer) reConnect() (<-chan amqp.Delivery, error) {
	be := backoff.NewExponentialBackOff()
	be.MaxElapsedTime = time.Minute
	be.InitialInterval = 1 * time.Second
	be.Multiplier = 2
	be.MaxInterval = c.maxInterval * time.Second

	b := backoff.WithContext(be, context.Background())
	for {
		d := b.NextBackOff()
		if d == backoff.Stop {
			return nil, errors.Errorf("stop reconnecting")
		}

		//nolint:gosimple
		select {
		case <-time.After(d):
			if err := c.connect(); err != nil {
				logrus.Printf("could not connect in reconnect call: %+v", err)
				continue
			}
			msgs, err := c.announceQueue()
			if err != nil {
				logrus.Printf("couldn't connect: %+v", err)
				continue
			}

			return msgs, nil
		}
	}
}

func (c *Consumer) Handle(fn func(<-chan amqp.Delivery)) error {
	var err error
	if err = c.connect(); err != nil {
		return errors.Wrap(err, "handle error")
	}
	msgs, err := c.announceQueue()
	if err != nil {
		return err
	}

	for {
		for i := 0; i < runtime.NumCPU(); i++ {
			go fn(msgs)
		}

		if <-c.done != nil {
			msgs, err = c.reConnect()
			if err != nil {
				return errors.Wrap(err, "reconnect error")
			}
		}
		logrus.Println("Reconnected!")
	}
}
