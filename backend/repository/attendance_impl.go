package repository

import (
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"gorm.io/gorm"
)

type AttendanceRepositoryImpl struct {
	db *gorm.DB
}

func NewAttendanceepositoryImpl(db *gorm.DB) *AttendanceRepositoryImpl {
	return &AttendanceRepositoryImpl{
		db: db,
	}
}

func (es AttendanceRepositoryImpl) InsertAttendance(ctx echo.Context, eventID openapi_types.UUID) error {
	// todo
	return nil
}

func (es AttendanceRepositoryImpl) DeleteAttendance(ctx echo.Context, eventID openapi_types.UUID) error {
	return nil
}
