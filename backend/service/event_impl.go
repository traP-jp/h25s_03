package service

import (
	"github.com/eraxyso/go-template/api"
	"github.com/eraxyso/go-template/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type EventServiceImpl struct {
	eventRepository    repository.EventRepository
	adminRepository    repository.AdminRepository
	attendeeRepository repository.AttendeeRepository
}

func NewEventServiceImpl(eventRepository repository.EventRepository, adminRepository repository.AdminRepository, attendeeRepository repository.AttendeeRepository) *EventServiceImpl {
	return &EventServiceImpl{
		eventRepository:    eventRepository,
		adminRepository:    adminRepository,
		attendeeRepository: attendeeRepository,
	}
}

func (es EventServiceImpl) PostEvents(ctx echo.Context, newEvent api.PostEventJSONRequestBody) (uuid.UUID, error) {
	// to use transaction
	eventID, err := es.eventRepository.InsertEvent(ctx, newEvent)
	if err != nil {
		return uuid.UUID{}, err
	}
	es.adminRepository.InsertAdmins(ctx, eventID, newEvent.Admins)
	es.attendeeRepository.InsertAttendees(ctx, eventID, newEvent.Attendees)

	return eventID, nil
}
func (es EventServiceImpl) GetEvents(ctx echo.Context, params api.GetEventsParams, userID string) ([]api.EventSummary, error) {
	ifDeleted := params.IfDeleted
	eventSummaries, err := es.eventRepository.GetEventSummaries(ctx, ifDeleted, userID)
	if err != nil {
		return nil, err
	}
	return eventSummaries, nil
}

func (es EventServiceImpl) GetEvent(ctx echo.Context, eventID uuid.UUID, userID string) (api.Event, error) {
	event, err := es.eventRepository.GetEvent(ctx, eventID, userID)
	if err != nil {
		return api.Event{}, err
	}
	return event, nil
}

func (es EventServiceImpl) PatchEvent(ctx echo.Context, eventID uuid.UUID, eventModification api.PatchEventJSONRequestBody) error {
	if err := es.eventRepository.UpdateEvent(ctx, eventID, eventModification); err != nil {
		return err
	}
	return nil
}

func (es EventServiceImpl) DeleteEvent(ctx echo.Context, eventID uuid.UUID) error {
	if err := es.eventRepository.DeleteEvent(ctx, eventID); err != nil {
		return err
	}
	return nil
}
