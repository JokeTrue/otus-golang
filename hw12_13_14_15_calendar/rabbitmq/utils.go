package rabbitmq

import (
	"github.com/sirupsen/logrus"
)

func handleError(err error, msg string) {
	if err != nil {
		logrus.Fatalf("%s: %s", msg, err)
	}
}
