package calendarscheduler

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type config struct{}

var Conf *config

func init() {
	configPath := pflag.String("config", "", "path to scheduler config")
	pflag.Parse()

	viper.SetConfigFile(*configPath)
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		logrus.Fatal(err)
	}
}
