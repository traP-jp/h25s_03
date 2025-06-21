package repository

import (
	"github.com/eraxyso/go-template/api"
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