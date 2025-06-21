package service

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AttendanceService interface {
	PostAttendance(ctx echo.Context, eventID uuid.UUID, userID string) error
	DeleteAttendance(ctx echo.Context, eventID uuid.UUID, userID string) error
}
