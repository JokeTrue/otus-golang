package usecase

import (
	"context"
	"time"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/event"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/models"

	"github.com/google/uuid"
)

type EventUseCase struct {
	eventRepo event.Repository
}

func NewEventUseCase(eventRepo event.Repository) *EventUseCase {
	return &EventUseCase{eventRepo: eventRepo}
}

func (e EventUseCase) CreateEvent(ctx context.Context, userID int64, title, description string, startDate, endDate string, notifyInterval time.Duration) (uuid.UUID, error) {
	ev, err := models.NewEvent(uuid.New(), userID, title, description, startDate, endDate, notifyInterval)
	if err != nil {
		return uuid.UUID{}, err
	}
	return e.eventRepo.CreateEvent(ctx, ev)
}

func (e EventUseCase) RetrieveEvent(ctx context.Context, userID int64, eventID uuid.UUID) (*models.Event, error) {
	return e.eventRepo.RetrieveEvent(ctx, userID, eventID)
}

func (e EventUseCase) DeleteEvent(ctx context.Context, userID int64, eventID uuid.UUID) error {
	return e.eventRepo.DeleteEvent(ctx, userID, eventID)
}

func (e EventUseCase) UpdateEvent(ctx context.Context, userID int64, ev *models.Event, eventID uuid.UUID) error {
	return e.eventRepo.UpdateEvent(ctx, userID, ev, eventID)
}

func (e EventUseCase) GetEvents(ctx context.Context, userID int64, interval event.Interval, startDate string) ([]*models.Event, error) {
	layout := "2006-01-02"
	parsedStartDate, err := time.Parse(layout, startDate)
	if err != nil {
		return nil, err
	}
	endDate := interval.GetIntervalEnd(parsedStartDate)
	return e.eventRepo.GetEvents(ctx, userID, parsedStartDate, endDate)
}
