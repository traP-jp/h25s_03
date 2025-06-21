package repository

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/eraxyso/go-template/api"
)

type EventRepository interface {
	InsertEvent(ctx echo.Context, event api.PostEventsJSONRequestBody) (uuid.UUID, error)
	RequestEventsSummary(ctx echo.Context, isDelete bool) ([]api.EventSummary, error)
	UpdateEvent(ctx echo.Context, eventID uuid.UUID, event api.PatchEventJSONRequestBody) error
	RemoveEvent(ctx echo.Context, eventID uuid.UUID) error
}
