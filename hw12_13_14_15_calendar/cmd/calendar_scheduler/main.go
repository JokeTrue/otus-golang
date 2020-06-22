package main

import (
	scheduler "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/calendar_scheduler"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/rabbitmq"
	"github.com/sirupsen/logrus"
)

func main() {
	sender := rabbitmq.NewSender()
	s := scheduler.NewCalendarScheduler(sender)
	if err := s.Run(); err != nil {
		logrus.Fatalf("%s", err.Error())
	}
}
