package main

import (
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/config"
	"github.com/sirupsen/logrus"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/server"
)

func main() {
	app := server.NewApp()
	if err := app.Run(config.Conf.App.Host, config.Conf.App.Port); err != nil {
		logrus.Fatalf("%s", err.Error())
	}
}
