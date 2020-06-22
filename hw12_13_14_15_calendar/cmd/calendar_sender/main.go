package main

import (
	sender "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/calendar_sender"
	calendarsender "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/config/calendar_sender"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/rabbitmq"
	"github.com/sirupsen/logrus"
)

func main() {
	connector := rabbitmq.NewConnector(
		calendarsender.Conf.URL,
		calendarsender.Conf.Exchange,
		calendarsender.Conf.QueueName,
		calendarsender.Conf.QOS,
	)

	s := sender.NewCalendarSender(connector)
	if err := s.Run(); err != nil {
		logrus.Fatalf("%s", err.Error())
	}
}
