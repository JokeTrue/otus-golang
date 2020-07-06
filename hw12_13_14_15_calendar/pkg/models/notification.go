package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	EventID  uuid.UUID
	UserID   int64
	Title    string
	Datetime time.Time
}
