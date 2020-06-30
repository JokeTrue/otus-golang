package event

import (
	"time"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/models"
	"github.com/google/uuid"
)

type UseCase interface {
	CreateEvent(
		userID int64,
		title string,
		description string,
		startDate, endDate string,
		notifyInterval time.Duration,
	) (uuid.UUID, error)
	RetrieveEvent(userID int64, eventID uuid.UUID) (*models.Event, error)
	UpdateEvent(userID int64, ev *models.Event, eventID uuid.UUID) error
	DeleteEvent(userID int64, eventID uuid.UUID) error
	GetEvents(userID int64, interval Interval, startDate string) ([]*models.Event, error)
}
