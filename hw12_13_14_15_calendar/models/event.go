package models

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/google/uuid"
)

type Event struct {
	ID             uuid.UUID     `gorm:"type:uuid;primary_key;" json:"id"`
	UserID         int64         `json:"user_id"`
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	StartDate      time.Time     `json:"start_date"`
	EndDate        time.Time     `json:"end_date"`
	NotifyInterval time.Duration `json:"notify_interval"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (e *Event) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.New())
}

func NewEvent(ID uuid.UUID, userID int64, title string, description string, startDate, endDate string, notifyInterval time.Duration) (*Event, error) {
	layout := "2006-01-02T15:04:05"
	parsedStartDate, err := time.Parse(layout, startDate)
	if err != nil {
		return nil, err
	}
	parsedEndDate, err := time.Parse(layout, endDate)
	if err != nil {
		return nil, err
	}

	return &Event{
		ID:             ID,
		UserID:         userID,
		Title:          title,
		Description:    description,
		StartDate:      parsedStartDate,
		EndDate:        parsedEndDate,
		NotifyInterval: notifyInterval,
	}, nil
}

func (e Event) GetDuration() time.Duration {
	return e.EndDate.Sub(e.StartDate)
}
