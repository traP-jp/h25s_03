package repository

import (
	"github.com/eraxyso/go-template/api"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type EventRepository interface {
	InsertEvent(ctx echo.Context, event api.PostEventsJSONRequestBody) error
	RequestEventsSummary(ctx echo.Context, isDelete bool) ([]api.EventSummary, error)
	DeleteEvent(ctx echo.Context, eventID openapi_types.UUID) error
	UpdateEvent(ctx echo.Context, eventID openapi_types.UUID, requestBody api.PatchEventJSONRequestBody) error

