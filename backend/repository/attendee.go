package repository

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AttendeeRepository interface {
	InsertAttendees(ctx echo.Context, eventID uuid.UUID, userIDs []string) error
	DeleteAttendees(ctx echo.Context, eventID uuid.UUID, userIDs []string) error
}
