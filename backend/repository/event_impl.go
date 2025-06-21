package repository

import (
	"time"

	"github.com/eraxyso/go-template/api"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type EventRepositoryImpl struct {
	db *gorm.DB
}

func NewEventRepositoryImpl(db *gorm.DB) *EventRepositoryImpl {
	return &EventRepositoryImpl{
		db: db,
	}
}

func (es EventRepositoryImpl) InsertEvent(ctx echo.Context, event api.PostEventsJSONRequestBody) error {
	// todo
	return nil
}

func (es EventRepositoryImpl) RequestEventsSummary(ctx echo.Context, isDelete bool) ([]api.EventSummary, error) {
	// todo
	return nil, nil
}

func (es EventRepositoryImpl) DeleteEvent(ctx echo.Context, eventID openapi_types.UUID) error{
	return nil
}

func (es EventRepositoryImpl) UptateEvent(ctx echo.Context, eventID openapi_types.UUID) error{
	return nil
}

type SelectEvent struct {
	EventId     uuid.UUID `json:"event_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Admins      []string  `json:"admins"`
	IsOpen      bool      `json:"is_open"`
	Attendees   []string  `json:"attendees"`
}

func (es EventRepositoryImpl) SelectEvent(ctx echo.Context, eventID uuid.UUID) (event SelectEvent, err error) {
	return SelectEvent{}, nil
}

// SELECT events.eventid, events.title, events.description, 
// events.data, events.isOpen, admins.userid, attendees.userid 
// FROM events JOIN admins ON events.eventid == admins.eventid 
// JOIN attendees ON events.eventid == attendees.eventid
