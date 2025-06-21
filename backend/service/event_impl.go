package service

import (
	"context"
	"fmt"
	"time"

	"github.com/eraxyso/go-template/repository"
	"github.com/google/uuid"
)

type EventServiceImpl struct {
	eventRepository    repository.EventRepository
	adminRepository    repository.AdminRepository
	attendeeRepository repository.AttendeeRepository
}

var _ EventService = &EventServiceImpl{}

func NewEventServiceImpl(eventRepository repository.EventRepository, adminRepository repository.AdminRepository, attendeeRepository repository.AttendeeRepository) *EventServiceImpl {
	return &EventServiceImpl{
		eventRepository:    eventRepository,
		adminRepository:    adminRepository,
		attendeeRepository: attendeeRepository,
	}
}

type EventOnCreate struct {
	Title       string
	Description string
	Date        time.Time
	IsOpen      bool
	Admins      []string
	Attendees   []string
}

func (es *EventServiceImpl) CreateEvent(ctx context.Context, event EventOnCreate) (uuid.UUID, error) {
	// TODO: use transaction
	newEvent := repository.EventOnCreate{
		Title:       event.Title,
		Description: event.Description,
		Date:        event.Date,
		IsOpen:      event.IsOpen,
	}
	eventID, err := es.eventRepository.InsertEvent(ctx, newEvent)
	if err != nil {
		return uuid.Nil, fmt.Errorf("insert event (service): %w", err)
	}
	err = es.adminRepository.InsertAdmins(ctx, eventID, event.Admins)
	if err != nil {
		return uuid.Nil, fmt.Errorf("insert admins (service): %w", err)
	}
	err = es.attendeeRepository.InsertAttendees(ctx, eventID, event.Attendees)
	if err != nil {
		return uuid.Nil, fmt.Errorf("insert attendees (service): %w", err)
	}
	return eventID, nil
}

type EventSummary struct {
	EventID      uuid.UUID
	Title        string
	Description  string
	Date         time.Time
	IsOpen       bool
	IsMeAttendee bool
	Admins       []string
}

func (es *EventServiceImpl) GetEvents(ctx context.Context, ifDeleted bool, userID string) ([]EventSummary, error) {
	eventWithAdminsAndAttendees, err := es.eventRepository.GetEvents(ctx, ifDeleted)
	if err != nil {
		return nil, fmt.Errorf("get events (service): %w", err)
	}
	eventSummaries := make([]EventSummary, 0, len(eventWithAdminsAndAttendees))
	for _, event := range eventWithAdminsAndAttendees {
		isMeAttendee := false
		if userID != "" {
			for _, attendee := range event.Attendees {
				if attendee == userID {
					isMeAttendee = true
					break
				}
			}
		}
		eventSummaries = append(eventSummaries, EventSummary{
			EventID:      event.EventID,
			Title:        event.Title,
			Description:  event.Description,
			Date:         event.Date,
			IsOpen:       event.IsOpen,
			Admins:       event.Admins,
			IsMeAttendee: isMeAttendee,
		})
	}
	return eventSummaries, nil
}

type EventDetail struct {
	EventID      uuid.UUID
	Title        string
	Description  string
	Date         time.Time
	IsOpen       bool
	IsMeAttendee bool
	IsDeleted    bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Admins       []string
	Attendees    []string
}

func (es *EventServiceImpl) GetEvent(ctx context.Context, eventID uuid.UUID, userID string) (EventDetail, error) {
	event, err := es.eventRepository.GetEvent(ctx, eventID)
	if err != nil {
		return EventDetail{}, fmt.Errorf("get event (service): %w", err)
	}
	isMeAttendee := false
	if userID != "" {
		for _, attendee := range event.Attendees {
			if attendee == userID {
				isMeAttendee = true
				break
			}
		}
	}
	detail := EventDetail{
		EventID:      event.EventID,
		Title:        event.Title,
		Description:  event.Description,
		Date:         event.Date,
		IsOpen:       event.IsOpen,
		IsMeAttendee: isMeAttendee,
		IsDeleted:    event.IsDeleted,
		CreatedAt:    event.CreatedAt,
		UpdatedAt:    event.UpdatedAt,
		Admins:       event.Admins,
		Attendees:    event.Attendees,
	}
	return detail, nil
}

type EventOnEdit struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	IsOpen      bool      `json:"is_open"`
	Admins      []string  `json:"admins"`
	Attendees   []string  `json:"attendees"`
}

func (es *EventServiceImpl) EditEvent(ctx context.Context, eventID uuid.UUID, event EventOnEdit) error {
	updateEvent := repository.EventOnUpdate{
		Title:       event.Title,
		Description: event.Description,
		Date:        event.Date,
		IsOpen:      event.IsOpen,
	}
	if err := es.eventRepository.UpdateEvent(ctx, eventID, updateEvent); err != nil {
		return fmt.Errorf("update event (service): %w", err)
	}
	if err := es.adminRepository.UpdateAdmins(ctx, eventID, event.Admins); err != nil {
		return fmt.Errorf("update admins (service): %w", err)
	}
	if err := es.attendeeRepository.UpdateAttendees(ctx, eventID, event.Attendees); err != nil {
		return fmt.Errorf("update attendees (service): %w", err)
	}
	return nil
}

func (es *EventServiceImpl) DeleteEvent(ctx context.Context, eventID uuid.UUID) error {
	if err := es.eventRepository.DeleteEvent(ctx, eventID); err != nil {
		return fmt.Errorf("delete event (service): %w", err)
	}
	return nil
}
