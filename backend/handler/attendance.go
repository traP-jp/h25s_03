package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handler) PostAttendance(ctx echo.Context, eventID openapi_types.UUID) error {
	userID := h.MiddlewareService.GetUserID(ctx)
	if err := h.AttendanceService.PostAttendance(ctx.Request().Context(), eventID, userID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("post attendance (handler): %w", err))
	}
	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) DeleteAttendance(ctx echo.Context, eventID openapi_types.UUID) error {
	userID := h.MiddlewareService.GetUserID(ctx)
	if err := h.AttendanceService.DeleteAttendance(ctx.Request().Context(), eventID, userID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("delete attendance (handler): %w", err))
	}
	return ctx.NoContent(http.StatusOK)
}
