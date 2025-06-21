package service

import (
	"github.com/eraxyso/go-template/api"
	"github.com/eraxyso/go-template/repository"
	"github.com/labstack/echo/v4"
)

type EventServiceImpl struct {
	eventRepository repository.EventRepository
}

func NewEventServiceImpl(exampleRepository repository.EventRepository) *EventServiceImpl {
	return &EventServiceImpl{}
}

func (es EventServiceImpl) CreateEvent(ctx echo.Context, event api.PostEventsJSONRequestBody) error {
	if err := es.eventRepository.InsertEvent(ctx, event); err != nil {
		return err
	}
	return nil
}
