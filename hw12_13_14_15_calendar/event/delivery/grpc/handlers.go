package grpc

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/models"

	"github.com/golang/protobuf/ptypes/duration"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/event"
	grpc2 "github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/event/delivery/grpc/schema"
)

type EventsServer struct {
	useCase event.UseCase
}

func NewEventsServer(useCase event.UseCase) *EventsServer {
	return &EventsServer{useCase: useCase}
}

func (e EventsServer) CreateEvent(ctx context.Context, request *grpc2.CreateEventRequest) (*grpc2.CreateEventResponse, error) {
	res := &grpc2.CreateEventResponse{}

	eventID, err := e.useCase.CreateEvent(
		request.UserID,
		request.Event.Title,
		request.Event.Description,
		request.Event.StartDate,
		request.Event.EndDate,
		time.Duration(request.Event.NotifyInterval.Seconds)*time.Second,
	)
	if err != nil {
		logrus.Error(errors.Wrap(err, "failed to create event"))
		res.Err = &grpc2.Error{
			Code:    uint32(codes.Internal),
			Message: err.Error(),
		}
		return res, err
	}

	res.ID = eventID.String()
	return res, nil
}

func (e EventsServer) RetrieveEvent(ctx context.Context, request *grpc2.RetrieveEventRequest) (*grpc2.RetrieveEventResponse, error) {
	res := &grpc2.RetrieveEventResponse{}

	eventID, err := uuid.Parse(request.ID)
	if err != nil {
		logrus.Error(errors.Wrap(err, "failed to parse event id"))
		res.Err = &grpc2.Error{
			Code:    uint32(codes.Internal),
			Message: err.Error(),
		}
		return res, err
	}
	ev, err := e.useCase.RetrieveEvent(request.UserID, eventID)
	if err != nil {
		logrus.Error(errors.Wrap(err, "failed to retrieve event"))
		res.Err = &grpc2.Error{
			Code:    uint32(codes.Internal),
			Message: err.Error(),
		}
		return res, err
	}
	res.Event = &grpc2.Event{
		ID:             ev.ID.String(),
		UserID:         ev.UserID,
		Title:          ev.Title,
		Description:    ev.Description,
		StartDate:      ev.StartDate.Format("2006-01-02T15:04:05"),
		EndDate:        ev.EndDate.Format("2006-01-02T15:04:05"),
		NotifyInterval: &duration.Duration{Seconds: int64(ev.NotifyInterval)},
	}

	return res, nil
}

func (e EventsServer) UpdateEvent(ctx context.Context, request *grpc2.UpdateEventRequest) (*grpc2.Error, error) {
	res := &grpc2.Error{}

	eventID, err := uuid.Parse(request.ID)
	if err != nil {
		logrus.Error(errors.Wrap(err, "failed to parse event id"))
		res.Code = uint32(codes.Internal)
		res.Message = err.Error()
		return res, err
	}

	ev, err := models.NewEvent(
		eventID,
		request.UserID,
		request.Event.Title,
		request.Event.Description,
		request.Event.StartDate,
		request.Event.EndDate,
		time.Duration(request.Event.NotifyInterval.Seconds)*time.Second,
	)
	if err != nil {
		res.Code = uint32(codes.Internal)
		res.Message = err.Error()
		return res, err
	}

	err = e.useCase.UpdateEvent(request.UserID, ev, eventID)
	if err != nil {
		logrus.Error(errors.Wrap(err, "failed to update event"))
		res.Code = uint32(codes.Internal)
		res.Message = err.Error()
		return res, err
	}

	return res, nil
}

func (e EventsServer) DeleteEvent(ctx context.Context, request *grpc2.DeleteEventRequest) (*grpc2.Error, error) {
	res := &grpc2.Error{}

	eventID, err := uuid.Parse(request.ID)
	if err != nil {
		logrus.Error(errors.Wrap(err, "failed to parse event id"))
		res.Code = uint32(codes.Internal)
		res.Message = err.Error()
		return res, err
	}

	err = e.useCase.DeleteEvent(request.UserID, eventID)
	if err != nil {
		logrus.Error(errors.Wrap(err, "failed to delete event"))
		res.Code = uint32(codes.Internal)
		res.Message = err.Error()
		return res, err
	}
	return res, nil
}

func (e EventsServer) GetEvents(ctx context.Context, request *grpc2.GetEventsRequest) (*grpc2.GetEventsResponse, error) {
	res := &grpc2.GetEventsResponse{}

	evs, err := e.useCase.GetEvents(request.UserID, event.Interval(request.Interval), request.StartDate)
	if err != nil {
		logrus.Error(errors.Wrap(err, "failed to retrieve list of events"))
		res.Err = &grpc2.Error{
			Code:    uint32(codes.Internal),
			Message: err.Error(),
		}
		return res, err
	}

	grpcEvents := make([]*grpc2.Event, len(evs))
	for _, ev := range evs {
		grpcEv := &grpc2.Event{
			ID:             ev.ID.String(),
			UserID:         ev.UserID,
			Title:          ev.Title,
			Description:    ev.Description,
			StartDate:      ev.StartDate.Format("2006-01-02T15:04:05"),
			EndDate:        ev.EndDate.Format("2006-01-02T15:04:05"),
			NotifyInterval: &duration.Duration{Seconds: int64(ev.NotifyInterval)},
		}
		grpcEvents = append(grpcEvents, grpcEv)
	}

	res.Events = grpcEvents
	return res, nil
}
