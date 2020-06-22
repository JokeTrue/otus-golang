package calendarsender

import (
	"encoding/json"
	"time"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/models"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/rabbitmq"
	"github.com/sirupsen/logrus"
)

type CalendarSender struct {
	subscriber rabbitmq.Subscriber
}

func NewCalendarSender(subscriber rabbitmq.Subscriber) *CalendarSender {
	return &CalendarSender{subscriber: subscriber}
}

func (s *CalendarSender) Run() error {
	msgCh := s.subscriber.Subscribe()

	for msg := range msgCh {
		var notification models.Notification
		err := json.Unmarshal(msg.Body, &notification)
		if err != nil {
			logrus.Error(err)
		}

		eta := int(notification.Datetime.Sub(time.Now()).Minutes())
		logrus.Debugf("Event [%s] for User[%d] is gonna be soon\n", notification.EventID.String(), notification.UserID)
		logrus.Printf("Get ready for the '%s' in %d minutes.\n", notification.Title, eta)
	}

	return nil
}
