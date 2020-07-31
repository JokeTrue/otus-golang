package psql

import (
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/pkg/models"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx" // Setup PGX Driver
)

type EventRepository struct {
	db *sqlx.DB
}

func NewEventRepository(db *sqlx.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (e EventRepository) CreateEvent(ev *models.Event) (id uuid.UUID, err error) {
	sqlStatement := `INSERT INTO events VALUES ($1, $2, $3, $4, $5, $6, $7)`
	err = e.db.QueryRowx(
		sqlStatement,
		ev.ID.String(),
		ev.UserID,
		ev.Title,
		ev.Description,
		ev.StartDate,
		ev.EndDate,
		ev.NotifyInterval,
	).Err()
	return ev.ID, err
}

func (e EventRepository) RetrieveEvent(userID int64, eventID uuid.UUID) (*models.Event, error) {
	var ev models.Event
	sqlStatement := `SELECT * FROM events WHERE user_id = $1 AND id = $2 LIMIT 1`
	err := e.db.Get(&ev, sqlStatement, userID, eventID)
	return &ev, err
}

func (e EventRepository) UpdateEvent(userID int64, ev *models.Event, eventID uuid.UUID) error {
	oldEv, err := e.RetrieveEvent(userID, eventID)
	if err != nil {
		return err
	}

	oldEv.Title = ev.Title
	oldEv.Description = ev.Description
	oldEv.StartDate = ev.StartDate
	oldEv.EndDate = ev.EndDate
	oldEv.NotifyInterval = ev.NotifyInterval

	sqlStatement := `UPDATE events SET
                     title = $1,
                     description = $2,
                     start_date = $3,
                     end_date = $4,
                     notify_interval = $5
					 WHERE user_id = $6 AND id = $7`

	return e.db.QueryRowx(
		sqlStatement,
		oldEv.Title,
		oldEv.Description,
		oldEv.StartDate,
		oldEv.EndDate,
		oldEv.NotifyInterval,
		userID,
		eventID,
	).Err()
}

func (e EventRepository) DeleteEvent(userID int64, eventID uuid.UUID) error {
	sqlStatement := `DELETE FROM events WHERE user_id = $1 AND id = $2`
	return e.db.QueryRowx(sqlStatement, userID, eventID).Err()
}

func (e EventRepository) GetUserEvents(userID int64, startDate time.Time, endDate time.Time) ([]*models.Event, error) {
	events := make([]*models.Event, 0)
	sqlStatement := `SELECT * FROM events WHERE user_id = $1 AND start_date BETWEEN $2 AND $3`
	err := e.db.Select(
		&events,
		sqlStatement,
		userID,
		startDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	)
	return events, err
}

func (e EventRepository) GetEvents(startDate, endDate time.Time) ([]*models.Event, error) {
	events := make([]*models.Event, 0)
	sqlStatement := `SELECT * FROM events WHERE start_date BETWEEN $1 AND $2`
	err := e.db.Select(
		&events,
		sqlStatement,
		startDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	)
	return events, err
}
