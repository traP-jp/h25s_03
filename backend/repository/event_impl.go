package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventRepositoryImpl struct {
	db *gorm.DB
}

var _ EventRepository = &EventRepositoryImpl{}

func NewEventRepositoryImpl(db *gorm.DB) *EventRepositoryImpl {
	return &EventRepositoryImpl{
		db: db,
	}
}

type eventModel struct {
	EventID     string    `gorm:"column:event_id;type:char(36);primaryKey;not null" json:"event_id"`
	Title       string    `gorm:"column:title;type:varchar(100);not null" json:"title"`
	Description string    `gorm:"column:description;type:text;not null" json:"description"`
	Date        time.Time `gorm:"column:date;type:date;not null" json:"date"`
	IsOpen      bool      `gorm:"column:is_open;type:boolean;not null" json:"is_open"`
	IsDeleted   bool      `gorm:"column:is_deleted;type:boolean;not null" json:"is_deleted"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`

	// Associations
	Admins    []adminModel    `gorm:"foreignKey:EventID;references:EventID" json:"admins,omitempty"`
	Attendees []attendeeModel `gorm:"foreignKey:EventID;references:EventID" json:"attendees,omitempty"`
	Lotteries []lotteryModel  `gorm:"foreignKey:EventID;references:EventID" json:"lotteries,omitempty"`
}

type EventOnCreate struct {
	Title       string
	Description string
	Date        time.Time
	IsOpen      bool
}

func (er *EventRepositoryImpl) InsertEvent(ctx context.Context, event EventOnCreate) (uuid.UUID, error) {
	eventID := uuid.New()
	newEvent := eventModel{
		EventID:     eventID.String(),
		Title:       event.Title,
		Description: event.Description,
		Date:        event.Date,
		IsOpen:      event.IsOpen,
		IsDeleted:   false,
	}
	if err := er.db.WithContext(ctx).Create(&newEvent).Error; err != nil {
		return uuid.Nil, fmt.Errorf("insert event (repository): %w", err)
	}
	return eventID, nil
}

type EventWithAdminsAndAttendees struct {
	EventID     uuid.UUID
	Title       string
	Description string
	Date        time.Time
	IsOpen      bool
	IsDeleted   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Admins      []string
	Attendees   []string
}

func (er *EventRepositoryImpl) GetEvents(ctx context.Context, ifDeleted bool) ([]EventWithAdminsAndAttendees, error) {
	query := er.db.WithContext(ctx).Preload("Admins").Preload("Attendees")
	if !ifDeleted {
		query = query.Where("is_deleted = ?", false)
	}
	var events []eventModel
	err := query.Find(&events).Error
	if err != nil {
		return nil, fmt.Errorf("get events (repository): %w", err)
	}

	var eventResult []EventWithAdminsAndAttendees
	for _, event := range events {
		eventID, err := uuid.Parse(event.EventID)
		if err != nil {
			return nil, fmt.Errorf("parse event id (repository): %w", err)
		}
		var admins []string
		for _, admin := range event.Admins {
			admins = append(admins, admin.TraqID)
		}
		var attendees []string
		for _, attendee := range event.Attendees {
			attendees = append(attendees, attendee.TraqID)
		}
		eventResult = append(eventResult, EventWithAdminsAndAttendees{
			EventID:     eventID,
			Title:       event.Title,
			Description: event.Description,
			Date:        event.Date,
			IsOpen:      event.IsOpen,
			IsDeleted:   event.IsDeleted,
			CreatedAt:   event.CreatedAt,
			UpdatedAt:   event.UpdatedAt,
			Admins:      admins,
			Attendees:   attendees,
		})
	}
	return eventResult, nil
}

func (er *EventRepositoryImpl) GetEvent(ctx context.Context, eventID uuid.UUID) (EventWithAdminsAndAttendees, error) {
	var event eventModel
	if err := er.db.WithContext(ctx).Preload("Admins").Preload("Attendees").Where("event_id = ?", eventID).Find(&event).Error; err != nil {
		return EventWithAdminsAndAttendees{}, fmt.Errorf("get event (repository): %w", err)
	}
	eventID, err := uuid.Parse(event.EventID)
	if err != nil {
		return EventWithAdminsAndAttendees{}, fmt.Errorf("parse event id (repository): %w", err)
	}
	var admins, attendees []string
	for _, admin := range event.Admins {
		admins = append(admins, admin.TraqID)
	}
	eventResult := EventWithAdminsAndAttendees{
		EventID:     eventID,
		Title:       event.Title,
		Description: event.Description,
		Date:        event.Date,
		IsOpen:      event.IsOpen,
		IsDeleted:   event.IsDeleted,
		CreatedAt:   event.CreatedAt,
		UpdatedAt:   event.UpdatedAt,
		Admins:      admins,
		Attendees:   attendees,
	}
	return eventResult, nil
}

type EventOnUpdate struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	IsOpen      bool      `json:"is_open,omitempty"`
}

func (er *EventRepositoryImpl) UpdateEvent(ctx context.Context, eventID uuid.UUID, event EventOnUpdate) error {
	eventOnUpdate := eventModel{
		Title:       event.Title,
		Description: event.Description,
		Date:        event.Date,
		IsOpen:      event.IsOpen,
	}
	if err := er.db.WithContext(ctx).Model(&eventModel{}).Where("event_id = ?", eventID.String()).Updates(eventOnUpdate).Error; err != nil {
		return fmt.Errorf("update event (repository): %w", err)
	}
	return nil
}

func (er *EventRepositoryImpl) DeleteEvent(ctx context.Context, eventID uuid.UUID) error {
	if err := er.db.WithContext(ctx).Model(&eventModel{}).Where("event_id = ?", eventID.String()).Update("is_deleted", true).Error; err != nil {
		return fmt.Errorf("delete event (repository): %w", err)
	}
	return nil
}
