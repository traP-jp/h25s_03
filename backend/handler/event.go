package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/eraxyso/go-template/api"
)

func (h *Handler) PostEvents(ctx echo.Context) error {
	requestBody := api.PostEventsJSONRequestBody{}
	if err := ctx.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if err := h.EventService.CreateEvent(ctx, requestBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return ctx.NoContent(http.StatusOK)
}
