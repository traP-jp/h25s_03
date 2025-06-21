package repository

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AdminRepository interface {
	InsertAdmins(ctx echo.Context, eventID uuid.UUID, userIDs []string) error
	RemoveAllAdmins(ctx echo.Context, eventID uuid.UUID) error
}
