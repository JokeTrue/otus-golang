package calendar

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	App struct {
		Host string
		Port string
		Type string
	}
	GRPC struct {
		Host string
		Port string
	}
	Logging struct {
		Path     string
		LogLevel string
	}
}

var Conf *config

func init() {
	configPath := pflag.String("config", "", "path to calendar config")
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
