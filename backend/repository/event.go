package repository

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/eraxyso/go-template/api"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type EventRepository interface {
	InsertEvent(ctx echo.Context, event api.PostEventsJSONRequestBody) (uuid.UUID, error)
	RequestEventsSummary(ctx echo.Context, isDelete bool) ([]api.EventSummary, error)
	DeleteEvent(ctx echo.Context, eventID openapi_types.UUID) error
	UpdateEvent(ctx echo.Context, eventID openapi_types.UUID, requestBody api.PatchEventJSONRequestBody) error
	SelectEvent(ctx echo.Context, eventID uuid.UUID) (SelectEvent, error)
	RemoveEvent(ctx echo.Context, eventID uuid.UUID) error
}
