package main

import (
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/calendar"
	"github.com/sirupsen/logrus"
)

func main() {
	app := calendar.NewApp()
	if err := app.Run(); err != nil {
		logrus.Fatalf("%s", err.Error())
	}
}
