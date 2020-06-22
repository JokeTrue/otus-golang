package main

import (
	scheduler "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/calendar_scheduler"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/event"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/event/usecase"
	config "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/internal/config/calendar_scheduler"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/internal/database"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/internal/utils"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/rabbitmq"
	"github.com/sirupsen/logrus"
)

func main() {
	connector := rabbitmq.NewConnector(
		config.Conf.Scheduler.URL,
		config.Conf.Scheduler.Exchange,
		config.Conf.Scheduler.QueueName,
		config.Conf.Scheduler.QOS,
	)
	s := scheduler.NewCalendarScheduler(connector)
	dsn := utils.GetDSN(
		config.Conf.Database.Host,
		config.Conf.Database.Port,
		config.Conf.Database.User,
		config.Conf.Database.Password,
		config.Conf.Database.Name,
	)
	s.EventUC = usecase.NewEventUseCase(event.GetEventRepository(database.GetDatabase(dsn)))

	if err := s.Run(); err != nil {
		logrus.Fatalf("%s", err.Error())
	}
}
