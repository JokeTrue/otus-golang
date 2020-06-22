package rabbitmq

type AMQPSender interface {
}

type Sender struct {
}

func NewSender() *Sender {
	return &Sender{}
}
