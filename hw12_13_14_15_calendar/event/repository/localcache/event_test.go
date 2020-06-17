package localcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/event"

	"github.com/stretchr/testify/assert"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/models"
	"github.com/google/uuid"
)

func createEvent(t *testing.T) *models.Event {
	ev, err := models.NewEvent(
		uuid.New(),
		1,
		"Встреча #1",
		"Встреча на метро Аэропорт",
		"2020-06-05T10:05:00",
		"2020-06-05T14:05:00",
		time.Duration(3600),
	)
	assert.NoError(t, err)
	return ev
}

func TestDeleteEvent(t *testing.T) {
	ev1 := createEvent(t)

	s := NewEventLocalStorage()
	_, err := s.CreateEvent(ev1)
	assert.NoError(t, err)

	err = s.DeleteEvent(1, ev1.ID)
	assert.NoError(t, err)

	err = s.DeleteEvent(1, ev1.ID)
	assert.Equal(t, err, event.ErrEventNotFound)
}

func TestCreateEvent(t *testing.T) {
	ev1 := createEvent(t)
	s := NewEventLocalStorage()
	id, err := s.CreateEvent(ev1)
	assert.NoError(t, err)
	assert.Equal(t, ev1.ID, id)
}

func TestRetrieveEvent(t *testing.T) {
	ev1 := createEvent(t)

	s := NewEventLocalStorage()

	retrieved, err := s.RetrieveEvent(1, uuid.New())
	assert.Equal(t, event.ErrEventNotFound, err)

	id, err := s.CreateEvent(ev1)
	assert.NoError(t, err)

	retrieved, err = s.RetrieveEvent(1, id)
	assert.Equal(t, ev1, retrieved)
}

func TestGetEvents(t *testing.T) {
	s := NewEventLocalStorage()

	cases := []struct {
		name               string
		interval           event.Interval
		startDate, endDate string
	}{
		{name: "day interval", interval: event.DayInterval, startDate: "2020-06-01", endDate: "2020-06-02"},
		{name: "week interval", interval: event.WeekInterval, startDate: "2020-06-05", endDate: "2020-06-06"},
		{name: "month interval", interval: event.MonthInterval, startDate: "2020-06-25", endDate: "2020-06-28"},
	}

	// Create Events
	createdEvents := make([]*models.Event, 0)
	for _, c := range cases {
		ev := createEvent(t)
		parsedDT, err := time.Parse("2006-01-02", c.startDate)
		assert.NoError(t, err)
		ev.StartDate = parsedDT

		_, err = s.CreateEvent(ev)
		assert.NoError(t, err)

		createdEvents = append(createdEvents, ev)
	}

	for i, c := range cases {
		startDate, err := time.Parse("2006-01-02", "2020-06-01")
		assert.NoError(t, err)
		endDate, err := time.Parse("2006-01-02", c.endDate)
		assert.NoError(t, err)

		list, err := s.GetEvents(1, startDate, endDate)
		assert.NoError(t, err)
		require.Len(t, list, i+1)
	}
}
