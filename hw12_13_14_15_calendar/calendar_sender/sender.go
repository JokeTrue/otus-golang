package calendarsender

import "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/rabbitmq"

type CalendarSender struct {
	consumer rabbitmq.AMQPConsumer
}

func NewCalendarSender(consumer rabbitmq.AMQPConsumer) *CalendarSender {
	return &CalendarSender{consumer: consumer}
}

func (s *CalendarSender) Run() error {
	return nil
}
