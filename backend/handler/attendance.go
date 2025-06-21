package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handler) DeleteAttendance(ctx echo.Context, eventID openapi_types.UUID) error {
	if err := h.AttendanceService.CancelAttendance(ctx, eventID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) PostAttendance(ctx echo.Context, eventID openapi_types.UUID) error{
	if err := h.AttendanceService.CreateAttendance(ctx, eventID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return ctx.NoContent(http.StatusOK)
}
