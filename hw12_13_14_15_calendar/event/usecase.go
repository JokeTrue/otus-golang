package event

import (
	"context"
	"time"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/models"
	"github.com/google/uuid"
)

type UseCase interface {
	CreateEvent(
		ctx context.Context,
		userID int64, title, description string,
		startDate, endDate string,
		notifyInterval time.Duration,
	) (uuid.UUID, error)
	RetrieveEvent(ctx context.Context, userID int64, eventID uuid.UUID) (*models.Event, error)
	UpdateEvent(ctx context.Context, userID int64, ev *models.Event, eventID uuid.UUID) error
	DeleteEvent(ctx context.Context, userID int64, eventID uuid.UUID) error
	GetEvents(ctx context.Context, userID int64, interval Interval, startDate string) ([]*models.Event, error)
}
