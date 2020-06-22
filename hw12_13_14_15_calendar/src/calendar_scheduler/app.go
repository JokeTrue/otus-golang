package calendarscheduler

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/event"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/models"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/src/rabbitmq"
)

type CalendarScheduler struct {
	publisher rabbitmq.Publisher
	EventUC   event.UseCase
}

func NewCalendarScheduler(publisher rabbitmq.Publisher) *CalendarScheduler {
	return &CalendarScheduler{publisher: publisher}
}

func (s CalendarScheduler) getEvents() []*models.Event {
	//TODO
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
	t := time.Tick(time.Duration(10) * time.Second)
	for {
		select {
		case <-t:
			err := s.dropOldEvents()
			logrus.Error(err)
		default:
		}
	}
}
