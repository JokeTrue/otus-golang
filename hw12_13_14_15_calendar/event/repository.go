package event

import (
	"time"

	"github.com/jinzhu/now"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/models"

	"github.com/google/uuid"
)

type Interval int

const (
	DayInterval Interval = iota + 1
	WeekInterval
	MonthInterval
)

func (i Interval) GetIntervalEnd(startDate time.Time) time.Time {
	var endDate time.Time
	switch i {
	case DayInterval:
		endDate = startDate.AddDate(0, 0, 1)
	case WeekInterval:
		endDate = now.With(startDate).EndOfWeek()
	case MonthInterval:
		endDate = now.With(startDate).EndOfMonth()
	}
	return endDate
}

type Repository interface {
	CreateEvent(ev *models.Event) (uuid.UUID, error)
	RetrieveEvent(userID int64, eventID uuid.UUID) (*models.Event, error)
	UpdateEvent(userID int64, ev *models.Event, eventID uuid.UUID) error
	DeleteEvent(userID int64, eventID uuid.UUID) error
	GetEvents(userID int64, startDate time.Time, endDate time.Time) ([]*models.Event, error)
}
