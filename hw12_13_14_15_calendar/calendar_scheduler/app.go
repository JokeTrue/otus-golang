package calendarscheduler

import (
	"encoding/json"
	"os"
	"os/signal"
	"time"

	config "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/config/calendar_scheduler"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/database"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/event/usecase"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/utils"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/models"

	"github.com/jinzhu/now"

	"github.com/sirupsen/logrus"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/event"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/rabbitmq"
)

type CalendarScheduler struct {
	publisher rabbitmq.Publisher
	EventUC   event.UseCase
	doneCh    chan os.Signal
}

func NewCalendarScheduler(publisher rabbitmq.Publisher) *CalendarScheduler {
	doneCh := make(chan os.Signal, 1)
	s := &CalendarScheduler{publisher: publisher, doneCh: doneCh}
	dsn := utils.GetDSN(
		config.Conf.Database.Host,
		config.Conf.Database.Port,
		config.Conf.Database.User,
		config.Conf.Database.Password,
		config.Conf.Database.Name,
	)
	s.EventUC = usecase.NewEventUseCase(event.GetEventRepository(config.Conf.App.Type, database.GetDatabase(dsn)))
	return s
}
func (s CalendarScheduler) notifyEvents() error {
	events, err := s.EventUC.GetEvents(now.BeginningOfDay(), now.EndOfDay())
	if err != nil {
		return err
	}

	for _, ev := range events {
		if time.Now().After(ev.StartDate.Add(-time.Duration(ev.NotifyInterval) * time.Second)) {
			notification := &models.Notification{
				EventID:  ev.ID,
				UserID:   ev.UserID,
				Title:    ev.Title,
				Datetime: ev.StartDate,
			}
			msg, err := json.Marshal(notification)
			if err != nil {
				return err
			}
			err = s.publisher.Publish(msg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s CalendarScheduler) dropOldEvents() error {
	startDate := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC)
	endDate := time.Now().AddDate(-1, 0, 0)

	events, err := s.EventUC.GetEvents(startDate, endDate)
	if err != nil {
		return err
	}

	for _, ev := range events {
		err = s.EventUC.DeleteEvent(ev.UserID, ev.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *CalendarScheduler) Run() error {
	signal.Notify(s.doneCh, os.Interrupt, os.Interrupt)
	t := time.NewTicker(time.Duration(10) * time.Second)
	for {
		select {
		case <-s.doneCh:
			return nil
		case <-t.C:
			err := s.notifyEvents()
			logrus.Error(err)
			err = s.dropOldEvents()
			logrus.Error(err)
		}
	}
}
