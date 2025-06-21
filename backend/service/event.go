package service

import (
	"github.com/eraxyso/go-template/api"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type EventService interface {
	PostEvents(ctx echo.Context, newEvent api.PostEventJSONRequestBody) (uuid.UUID, error)
	GetEvents(ctx echo.Context, ifDeleted bool, userID string) ([]api.EventSummary, error)
	GetEvent(ctx echo.Context, eventID uuid.UUID, userID string) (api.Event, error)
	PatchEvent(ctx echo.Context, eventID uuid.UUID, eventModification api.PatchEventJSONRequestBody) error
	DeleteEvent(ctx echo.Context, eventID uuid.UUID) error
}
