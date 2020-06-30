package calendar

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	}
	App struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Type string `yaml:"type"`
	}
	GRPC struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	Logging struct {
		Path     string `yaml:"path"`
		LogLevel string `yaml:"loglevel"`
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
