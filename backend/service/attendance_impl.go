package service

import (
	"context"
	"fmt"

	"traquji/repository"

	"github.com/google/uuid"
)

type AttendanceServiceImpl struct {
	attendeeRepository repository.AttendeeRepository
}

var _ AttendanceService = &AttendanceServiceImpl{}

func NewAttendanceServiceImpl(attendeeRepository repository.AttendeeRepository) *AttendanceServiceImpl {
	return &AttendanceServiceImpl{
		attendeeRepository: attendeeRepository,
	}
}

func (as *AttendanceServiceImpl) PostAttendance(ctx context.Context, eventID uuid.UUID, userID string) error {
	if err := as.attendeeRepository.InsertAttendees(ctx, eventID, []string{userID}); err != nil {
		return fmt.Errorf("insert attendees (service): %w", err)
	}
	return nil
}

func (as *AttendanceServiceImpl) DeleteAttendance(ctx context.Context, eventID uuid.UUID, userID string) error {
	if err := as.attendeeRepository.DeleteAttendees(ctx, eventID, []string{userID}); err != nil {
		return fmt.Errorf("delete attendees (service): %w", err)
	}
	return nil
}
