package calendarscheduler

import "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/rabbitmq"

type CalendarScheduler struct {
	sender rabbitmq.AMQPSender
}

func NewCalendarScheduler(sender rabbitmq.AMQPSender) *CalendarScheduler {
	return &CalendarScheduler{sender: sender}
}

func (s *CalendarScheduler) Run() error {
	return nil
}
