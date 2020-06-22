package main

import (
	sender "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/calendar_sender"
	config "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/internal/config/calendar_sender"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/rabbitmq"
	"github.com/sirupsen/logrus"
)

func main() {
	connector := rabbitmq.NewConnector(
		config.Conf.URL,
		config.Conf.Exchange,
		config.Conf.QueueName,
		config.Conf.QOS,
	)

	s := sender.NewCalendarSender(connector)
	if err := s.Run(); err != nil {
		logrus.Fatalf("%s", err.Error())
	}
}
