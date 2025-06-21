package repository

import (
	"github.com/eraxyso/go-template/api"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type EventRepository interface {
	InsertEvent(ctx echo.Context, newEvent api.PostEventsJSONRequestBody) (uuid.UUID, error)
	GetEventSummaries(ctx echo.Context, ifDeleted bool, userID string) ([]api.EventSummary, error)
	GetEvent(ctx echo.Context, eventID uuid.UUID, userID string) (api.Event, error)
	UpdateEvent(ctx echo.Context, eventID uuid.UUID, eventModification api.PatchEventJSONRequestBody) error
	DeleteEvent(ctx echo.Context, eventID uuid.UUID) error
}
