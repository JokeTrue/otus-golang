package main

import (
	"log"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/config"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/server"
	"github.com/spf13/viper"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	host, port := viper.GetString("app.host"), viper.GetString("app.port")
	app := server.NewApp()
	if err := app.Run(host, port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
