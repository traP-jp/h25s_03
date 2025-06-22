package service

import (
	"context"

	"github.com/google/uuid"
)

type AttendanceService interface {
	PostAttendance(ctx context.Context, eventID uuid.UUID, userID string) error
	DeleteAttendance(ctx context.Context, eventID uuid.UUID, userID string) error
}
