package service

import (
	"github.com/eraxyso/go-template/api"
	"github.com/eraxyso/go-template/repository"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
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



func (es EventServiceImpl)DeleteEvent(ctx echo.Context, eventID openapi_types.UUID) error{
	if err := es.eventRepository.DeleteEvent(ctx, eventID); err != nil {
		return err
	}
	return nil
}

func (es EventServiceImpl)EditEvent(ctx echo.Context, eventID openapi_types.UUID,requestBody api.PatchEventJSONRequestBody) error{
	if err := es.eventRepository.UpdateEvent(ctx, eventID,requestBody); err != nil {
		return err
	}
	return nil
}