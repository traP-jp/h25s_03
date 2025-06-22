package service

import (
	"fmt"
	"net/http"

	"github.com/eraxyso/go-template/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type MiddlewareServiceImpl struct {
	adminRepository repository.AdminRepository
	eventRepository repository.EventRepository
}

func NewMiddlewareServiceImpl() *MiddlewareServiceImpl {
	return &MiddlewareServiceImpl{}
}

func (ms *MiddlewareServiceImpl) GetUserID(ctx echo.Context) string {
	userID := ctx.Request().Header.Get("X-Forwarded-User")
	if userID == "" {
		userID = "cp20"
	}
	return userID
}

func (ms *MiddlewareServiceImpl) EventAdminAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userID := ms.GetUserID(ctx)
		paramEventID := ctx.Param("eventID")
		eventID, err := uuid.Parse(paramEventID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("invalid event ID: %w", err))
		}
		isAdmin, err := ms.adminRepository.CheckAdmin(ctx.Request().Context(), eventID, userID)
		if !isAdmin {
			return echo.NewHTTPError(http.StatusForbidden, "you are not an admin of this event")
		}
		return next(ctx)
	}
}

func (ms *MiddlewareServiceImpl) EventRegistrationAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		paramEventID := ctx.Param("eventID")
		eventID, err := uuid.Parse(paramEventID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("invalid event ID: %w", err))
		}
		event, err := ms.eventRepository.GetEvent(ctx.Request().Context(), eventID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("failed to get event: %w", err))
		}
		if !event.IsOpen {
			return echo.NewHTTPError(http.StatusForbidden, "event registration is closed")
		}
		return next(ctx)
	}
}
