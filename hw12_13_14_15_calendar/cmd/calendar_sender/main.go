package main

import (
	sender "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/calendar_sender"
	config "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/internal/config/calendar_sender"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/rabbitmq"
	"github.com/sirupsen/logrus"
)

func main() {
	consumer := rabbitmq.NewConsumer(
		config.Conf.ConsumerTag,
		config.Conf.URI,
		config.Conf.ExchangeName,
		config.Conf.ExchangeType,
		config.Conf.Queue,
		config.Conf.BindingKey,
		config.Conf.MaxInterval,
	)
	s := sender.NewCalendarSender(consumer)
	if err := s.Run(); err != nil {
		logrus.Fatalf("%s", err.Error())
	}
}
