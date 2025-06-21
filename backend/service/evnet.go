package service

import (
	"github.com/eraxyso/go-template/api"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type EventService interface {
	CreateEvent(ctx echo.Context, event api.PostEventsJSONRequestBody) error

	DeleteEvent(ctx echo.Context, eventID openapi_types.UUID) error

	EditEvent(ctx echo.Context, eventID openapi_types.UUID, requestBody api.PatchEventJSONRequestBody) error
}
