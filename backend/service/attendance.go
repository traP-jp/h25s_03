package service

import (
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type AttendanceService interface {
	CancelAttendance(ctx echo.Context, eventID openapi_types.UUID) error
	CreateAttendance(ctx echo.Context, eventID openapi_types.UUID) error
}


