package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Publisher interface {
	Publish(body []byte) error
}

type Subscriber interface {
	Subscribe() <-chan amqp.Delivery
}

type AMQPConnector interface {
	Publisher
	Subscriber
}

type connector struct {
	exchange string
	conn     *amqp.Connection
	channel  *amqp.Channel
	queue    amqp.Queue
	QOS      int
	msgCh    <-chan amqp.Delivery
}

//nolint
func NewConnector(url, exchange, queueName string, qos int) *connector {
	c := &connector{exchange: exchange, QOS: qos}
	c.setupConn(url)
	c.declareQueue(queueName)
	return c
}

func (c *connector) Publish(body []byte) error {
	err := c.channel.Publish(c.exchange, c.queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})
	return err
}

func (c *connector) Subscribe() <-chan amqp.Delivery {
	c.setupQOS(c.QOS)
	c.setupMsgCh()
	return c.msgCh
}

func (c *connector) setupConn(url string) {
	var err error
	c.conn, err = amqp.Dial(url)
	handleError(err, "Can't connect to AMQP")

	c.channel, err = c.conn.Channel()
	handleError(err, "Can't create a amqpChannel")
}

func (c *connector) declareQueue(name string) {
	var err error
	c.queue, err = c.channel.QueueDeclare(name, true, false, false, false, nil)
	handleError(err, fmt.Sprintf("Could not declare `%s` queue", name))
}

func (c *connector) setupQOS(count int) {
	err := c.channel.Qos(count, 0, false)
	handleError(err, "Could not configure QoS")
}

func (c *connector) setupMsgCh() {
	var err error
	c.msgCh, err = c.channel.Consume(
		c.queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Could not register consumer")
}
