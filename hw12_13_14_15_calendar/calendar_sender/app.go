package calendarsender

import (
	"encoding/json"
	"os"
	"os/signal"
	"time"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/models"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/rabbitmq"
	"github.com/sirupsen/logrus"
)

type CalendarSender struct {
	subscriber rabbitmq.Subscriber
	doneCh     chan os.Signal
}

func NewCalendarSender(subscriber rabbitmq.Subscriber) *CalendarSender {
	doneCh := make(chan os.Signal, 1)
	return &CalendarSender{subscriber: subscriber, doneCh: doneCh}
}

func (s *CalendarSender) Run() error {
	signal.Notify(s.doneCh, os.Interrupt, os.Interrupt)
	msgCh := s.subscriber.Subscribe()

	for {
		select {
		case <-s.doneCh:
			return nil
		case msg := <-msgCh:
			var notification models.Notification
			err := json.Unmarshal(msg.Body, &notification)
			if err != nil {
				logrus.Error(err)
			}

			eta := time.Until(notification.Datetime)
			logrus.Debugf("Event [%s] for User[%d] is gonna be soon\n", notification.EventID.String(), notification.UserID)
			logrus.Printf("Get ready for the '%s' in %d minutes.\n", notification.Title, eta)
		}
	}
}
