package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	EventID  uuid.UUID
	UserID   int32
	Title    string
	Datetime time.Time
}
