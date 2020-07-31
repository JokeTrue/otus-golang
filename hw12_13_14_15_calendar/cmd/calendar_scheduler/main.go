package main

import (
	scheduler "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/pkg/calendar_scheduler"
	calendarscheduler "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/pkg/config/calendar_scheduler"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/pkg/rabbitmq"
	"github.com/sirupsen/logrus"
)

func main() {
	connector := rabbitmq.NewConnector(
		calendarscheduler.Conf.Scheduler.URL,
		calendarscheduler.Conf.Scheduler.Exchange,
		calendarscheduler.Conf.Scheduler.QueueName,
		calendarscheduler.Conf.Scheduler.QOS,
	)
	s := scheduler.NewCalendarScheduler(connector)
	if err := s.Run(); err != nil {
		logrus.Fatalf("%s", err.Error())
	}
}
