package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID             uuid.UUID `json:"id" db:"id"`
	UserID         int64     `json:"user_id" db:"user_id"`
	Title          string    `json:"title" db:"title"`
	Description    string    `json:"description" db:"description"`
	StartDate      time.Time `json:"start_date" db:"start_date"`
	EndDate        time.Time `json:"end_date" db:"end_date"`
	NotifyInterval float64   `json:"notify_interval" db:"notify_interval"`
}

const (
	DefaultDatetimeLayout = "2006-01-02T15:04:05"
	DefaultDateLayout     = "2006-01-02"
)

func NewEvent(ID uuid.UUID, userID int64, title string, description string, startDate, endDate string, notifyInterval time.Duration) (*Event, error) {
	parsedStartDate, err := time.Parse(DefaultDatetimeLayout, startDate)
	if err != nil {
		return nil, err
	}
	parsedEndDate, err := time.Parse(DefaultDatetimeLayout, endDate)
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
		NotifyInterval: notifyInterval.Seconds(),
	}, nil
}

func (e Event) GetDuration() time.Duration {
	return e.EndDate.Sub(e.StartDate)
}
