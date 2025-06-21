package repository

import (
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type AttendanceRepository interface {
	DeleteAttendance(ctx echo.Context, eventID openapi_types.UUID) error
	InsertAttendance(ctx echo.Context, eventID openapi_types.UUID) error
}
