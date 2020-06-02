package psql

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/models"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx" // Setup PGX Driver
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (e EventRepository) CreateEvent(ctx context.Context, ev *models.Event) (id uuid.UUID, err error) {
	if err = e.db.Create(ev).Error; err != nil {
		return
	}
	return ev.ID, nil
}

func (e EventRepository) RetrieveEvent(ctx context.Context, userID int64, eventID uuid.UUID) (*models.Event, error) {
	var ev models.Event
	return &ev, e.db.Where(&models.Event{ID: eventID, UserID: userID}).First(&ev).Error
}

func (e EventRepository) UpdateEvent(ctx context.Context, userID int64, ev *models.Event, eventID uuid.UUID) error {
	oldEv, err := e.RetrieveEvent(ctx, userID, eventID)
	if err != nil {
		return err
	}

	oldEv.StartDate = ev.StartDate
	oldEv.EndDate = ev.EndDate
	oldEv.NotifyInterval = ev.NotifyInterval
	oldEv.Title = ev.Title
	oldEv.Description = ev.Description

	return e.db.Save(oldEv).Error
}

func (e EventRepository) DeleteEvent(ctx context.Context, userID int64, eventID uuid.UUID) error {
	ev, err := e.RetrieveEvent(ctx, userID, eventID)
	if err != nil {
		return err
	}
	return e.db.Delete(ev).Error
}

func (e EventRepository) GetEvents(ctx context.Context, userID int64, startDate time.Time, endDate time.Time) ([]*models.Event, error) {
	var events []*models.Event
	if err := e.db.Where(
		"user_id = ? AND start_date BETWEEN ? AND ?",
		userID,
		startDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}
