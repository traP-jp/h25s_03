package service

import (
	"github.com/eraxyso/go-template/repository"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type AttendanceServiceImpl struct {
	attendanceRepository repository.AttendanceRepository
}

func NewAttendanceServiceImpl(exampleRepository repository.AttendanceRepository) *AttendanceServiceImpl {
	return &AttendanceServiceImpl{}
}

func (es AttendanceServiceImpl) CancelAttendance(ctx echo.Context, eventID openapi_types.UUID) error{
	if err := es.attendanceRepository.DeleteAttendance(ctx, eventID); err != nil {
		return err
	}
	return nil	
}

func (es AttendanceServiceImpl) CreateEvent(ctx echo.Context, eventID openapi_types.UUID) error {
	if err := es.attendanceRepository.InsertAttendance(ctx, eventID); err != nil {
		return err
	}
	return nil
}
