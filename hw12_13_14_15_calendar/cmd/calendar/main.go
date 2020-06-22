package main

import (
	config "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/internal/config/calendar"
	"github.com/sirupsen/logrus"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/calendar"
)

func main() {
	app := calendar.NewApp()
	if err := app.Run(config.Conf.App.Host, config.Conf.App.Port); err != nil {
		logrus.Fatalf("%s", err.Error())
	}
}
