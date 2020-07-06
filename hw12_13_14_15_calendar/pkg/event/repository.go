package event

import (
	"time"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/pkg/event/repository/localcache"
	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/pkg/event/repository/psql"
	"github.com/jmoiron/sqlx"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/pkg/models"

	"github.com/google/uuid"
)

type Repository interface {
	CreateEvent(ev *models.Event) (uuid.UUID, error)
	RetrieveEvent(userID int64, eventID uuid.UUID) (*models.Event, error)
	UpdateEvent(userID int64, ev *models.Event, eventID uuid.UUID) error
	DeleteEvent(userID int64, eventID uuid.UUID) error
	GetUserEvents(userID int64, startDate time.Time, endDate time.Time) ([]*models.Event, error)
	GetEvents(startDate, endDate time.Time) ([]*models.Event, error)
}

func GetEventRepository(appType string, db *sqlx.DB) (eventRepo Repository) {
	switch appType {
	case "local":
		eventRepo = localcache.NewEventLocalStorage()
	case "psql":
		eventRepo = psql.NewEventRepository(db)
	}
	return eventRepo
}
