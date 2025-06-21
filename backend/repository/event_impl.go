package repository

import (
	"time"

	"github.com/eraxyso/go-template/api"
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

func (es EventRepositoryImpl) InsertEvent(ctx echo.Context, event api.PostEventsJSONRequestBody) error {
	// todo
	return nil
}

func (es EventRepositoryImpl) RequestEventsSummary(ctx echo.Context, isDelete bool) ([]api.EventSummary, error) {
	// todo
	return nil, nil
}

func (es EventRepositoryImpl) DeleteEvent(ctx echo.Context, eventID openapi_types.UUID) error {
	return nil
}

func (es EventRepositoryImpl) UptateEvent(ctx echo.Context, eventID openapi_types.UUID) error {
	return nil
}
