package repository

import (
	"github.com/eraxyso/go-template/api"
	"github.com/labstack/echo/v4"
)

type EventRepository interface {
	InsertEvent(ctx echo.Context, event api.PostEventsJSONRequestBody) error
	RequestEventsSummary(ctx echo.Context, isDelete bool) ([]api.EventSummary, error)
}
