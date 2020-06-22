package main

import (
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/event"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/event/usecase"
	config "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/internal/config/calendar"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/internal/database"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/internal/utils"
	"github.com/sirupsen/logrus"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/calendar"
)

func main() {
	dsn := utils.GetDSN(
		config.Conf.Database.Host,
		config.Conf.Database.Port,
		config.Conf.Database.User,
		config.Conf.Database.Password,
		config.Conf.Database.Name,
	)
	app := calendar.NewApp(database.GetDatabase(dsn))
	app.EventUC = usecase.NewEventUseCase(event.GetEventRepository(app.DB))

	if err := app.Run(config.Conf.App.Host, config.Conf.App.Port); err != nil {
		logrus.Fatalf("%s", err.Error())
	}
}
