package localcache

import (
	"sync"
	"time"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/models"
	"github.com/google/uuid"
)

type EventLocalStorage struct {
	events map[uuid.UUID]*models.Event
	mutex  *sync.Mutex
}

func NewEventLocalStorage() *EventLocalStorage {
	return &EventLocalStorage{
		events: make(map[uuid.UUID]*models.Event),
		mutex:  new(sync.Mutex),
	}
}

func (e EventLocalStorage) CreateEvent(ev *models.Event) (uuid.UUID, error) {
	e.mutex.Lock()
	e.events[ev.ID] = ev
	e.mutex.Unlock()
	return ev.ID, nil
}

func (e EventLocalStorage) RetrieveEvent(userID int64, eventID uuid.UUID) (*models.Event, error) {
	ev, ok := e.events[eventID]
	if ok && ev.UserID == userID {
		return ev, nil
	}
	return nil, ErrEventNotFound
}

func (e EventLocalStorage) UpdateEvent(userID int64, ev *models.Event, eventID uuid.UUID) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	oldEv, ok := e.events[eventID]
	if ok && oldEv.UserID == userID {
		e.events[eventID] = ev
		return nil
	}

	return ErrEventNotFound
}

func (e EventLocalStorage) DeleteEvent(userID int64, eventID uuid.UUID) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	ev, ok := e.events[eventID]
	if ok && ev.UserID == userID {
		delete(e.events, eventID)
		return nil
	}

	return ErrEventNotFound
}

func (e EventLocalStorage) GetUserEvents(userID int64, startDate time.Time, endDate time.Time) ([]*models.Event, error) {
	events := make([]*models.Event, 0)

	e.mutex.Lock()
	for _, ev := range e.events {
		if ev.UserID == userID && (ev.StartDate.After(startDate) || ev.StartDate.Equal(startDate)) && ev.StartDate.Before(endDate) {
			events = append(events, ev)
		}
	}
	e.mutex.Unlock()

	return events, nil
}

func (e EventLocalStorage) GetEvents(startDate, endDate time.Time) ([]*models.Event, error) {
	events := make([]*models.Event, 0)

	e.mutex.Lock()
	for _, ev := range e.events {
		if (ev.StartDate.After(startDate) || ev.StartDate.Equal(startDate)) && ev.StartDate.Before(endDate) {
			events = append(events, ev)
		}
	}
	e.mutex.Unlock()

	return events, nil
}
