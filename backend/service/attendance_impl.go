package service

import (
	"github.com/eraxyso/go-template/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AttendanceServiceImpl struct {
	attendeeRepository repository.AttendeeRepository
}

func NewAttendanceServiceImpl(attendeeRepository repository.AttendeeRepository) *AttendanceServiceImpl {
	return &AttendanceServiceImpl{
		attendeeRepository: attendeeRepository,
	}
}

func (as AttendanceServiceImpl) PostAttendance(ctx echo.Context, eventID uuid.UUID, userID string) error {
	if err := as.attendeeRepository.InsertAttendees(ctx, eventID, []string{userID}); err != nil {
		return err
	}
	return nil
}

func (as AttendanceServiceImpl) DeleteAttendance(ctx echo.Context, eventID uuid.UUID, userID string) error {
	if err := as.attendeeRepository.DeleteAttendees(ctx, eventID, []string{userID}); err != nil {
		return err
	}
	return nil
}
