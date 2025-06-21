package repository

import (
	"github.com/eraxyso/go-template/api"
	"github.com/labstack/echo/v4"
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

func (e *EventRepositoryImpl) EventRepositoryMethod() error {
	return nil
}

func (es EventRepositoryImpl) InsertEvent(ctx echo.Context, event api.PostEventsJSONRequestBody) error {
	// todo
	return nil
}
