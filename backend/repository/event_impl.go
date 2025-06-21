package repository

import (
	"time"

	"github.com/eraxyso/go-template/api"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"gorm.io/gorm"
)

type EventRepositoryImpl struct {
	db *gorm.DB
}

func NewEventRepositoryImpl(db *gorm.DB) *EventRepositoryImpl {
	return &EventRepositoryImpl{
		db: db,
	}
}

type Event struct {
	EventID     string    `gorm:"column:event_id;type:char(36);primaryKey;not null" json:"event_id"`
	Title       string    `gorm:"column:title;type:varchar(100);not null" json:"title"`
	Description string    `gorm:"column:description;type:text;not null" json:"description"`
	Date        time.Time `gorm:"column:date;type:date;not null" json:"date"`
	IsOpen      bool      `gorm:"column:is_open;type:boolean;not null" json:"is_open"`
	IsDeleted   bool      `gorm:"column:is_deleted;type:boolean;not null" json:"is_deleted"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`

	// Associations
	Admins    []Admin    `gorm:"foreignKey:EventID;references:EventID" json:"admins,omitempty"`
	Attendees []Attendee `gorm:"foreignKey:EventID;references:EventID" json:"attendees,omitempty"`
	Lotteries []Lottery  `gorm:"foreignKey:EventID;references:EventID" json:"lotteries,omitempty"`
}

func (er EventRepositoryImpl) InsertEvent(ctx echo.Context, newEvent api.PostEventsJSONRequestBody) (uuid.UUID, error) {
	eventID := uuid.New()
	event := Event{
		EventID:     eventID.String(),
		Title:       newEvent.Title,
		Description: newEvent.Description,
		Date:        newEvent.Date.Time,
		IsOpen:      newEvent.IsOpen,
		IsDeleted:   false,
	}
	if err := er.db.WithContext(ctx.Request().Context()).Create(&event).Error; err != nil {
		return uuid.UUID{}, nil
	}
	return eventID, nil
}

func (er EventRepositoryImpl) GetEventSummaries(ctx echo.Context, ifDeleted bool, userID string) ([]api.EventSummary, error) {
	query := er.db.WithContext(ctx.Request().Context()).Preload("Admins")
	if !ifDeleted {
		query = query.Where("is_deleted = ?", false)
	}
	var events []Event
	err := query.Find(&events).Error
	if err != nil {
		return nil, err
	}

	var eventSammaries []api.EventSummary
	for _, event := range events {
		eventID, err := uuid.Parse(event.EventID)
		if err != nil {
			return nil, err
		}
		isMeAttendee := false
		for _, attendee := range event.Attendees {
			if attendee.TraqID == userID {
				isMeAttendee = true
				break
			}
		}
		var admins []string
		for _, admin := range event.Admins {
			admins = append(admins, admin.TraqID)
		}
		eventSammaries = append(eventSammaries, api.EventSummary{
			EventId:      eventID,
			Title:        event.Title,
			Description:  event.Description,
			Date:         openapi_types.Date{Time: event.Date},
			IsOpen:       event.IsOpen,
			IsMeAttendee: isMeAttendee,
			Admins:       admins,
		})
	}
	return eventSammaries, nil
}

func (er EventRepositoryImpl) GetEvent(ctx echo.Context, eventID uuid.UUID, userID string) (api.Event, error) {
	var event Event
	if err := er.db.WithContext(ctx.Request().Context()).Preload("Admins").Preload("Attendees").Where("event_id = ?", eventID).Find(&event).Error; err != nil {
		return api.Event{}, err
	}
	eventID, err := uuid.Parse(event.EventID)
	if err != nil {
		return api.Event{}, err
	}
	var admins, attendees []string
	for _, admin := range event.Admins {
		admins = append(admins, admin.TraqID)
	}
	isMeAttendee := false
	for _, attendee := range event.Attendees {
		if attendee.TraqID == userID {
			isMeAttendee = true
		}
		attendees = append(attendees, attendee.TraqID)
	}
	apiEvent := api.Event{
		EventId:      eventID,
		Title:        event.Title,
		Description:  event.Description,
		Date:         openapi_types.Date{Time: event.Date},
		IsOpen:       event.IsOpen,
		IsDeleted:    event.IsDeleted,
		IsMeAttendee: isMeAttendee,
		CreatedAt:    event.CreatedAt,
		UpdatedAt:    event.UpdatedAt,
		Admins:       admins,
		Attendees:    attendees,
	}
	return apiEvent, nil
}

func (er EventRepositoryImpl) UpdateEvent(ctx echo.Context, eventID uuid.UUID, eventModification api.PatchEventJSONRequestBody) error {

	eventUpdate := make(map[string]interface{})
	eventUpdate["Title"] = eventModification.Title
	eventUpdate["Description"] = eventModification.Description
	eventUpdate["Date"] = eventModification.Date.Time
	eventUpdate["IsOpen"] = eventModification.IsOpen

	if err := er.db.WithContext(ctx.Request().Context()).Model(&Event{}).Where("event_id = ?", eventID.String()).Updates(eventUpdate).Error; err != nil {
		return err
	}
	return nil
}

func (er EventRepositoryImpl) DeleteEvent(ctx echo.Context, eventID uuid.UUID) error {
	if err := er.db.WithContext(ctx.Request().Context()).Model(&Event{}).Where("event_id = ?", eventID.String()).Update("is_deleted", true).Error; err != nil {
		return err
	}
	return nil
}
