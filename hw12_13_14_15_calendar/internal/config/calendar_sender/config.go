package calendarsender

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type config struct {
	ConsumerTag  string
	URI          string
	ExchangeName string
	ExchangeType string
	Queue        string
	BindingKey   string
	MaxInterval  time.Duration
}

var Conf *config

func init() {
	configPath := pflag.String("config", "", "path to sender config")
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
