package service

import (
	"github.com/eraxyso/go-template/api"
	"github.com/labstack/echo/v4"
)

type EventService interface {
	CreateEvent(ctx echo.Context, event api.PostEventsJSONRequestBody) error
	GetEventsSummary(ctx echo.Context, isDelete bool) ([]api.EventSummary, error)
}
