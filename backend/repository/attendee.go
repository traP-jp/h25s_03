package repository

import (
	"context"

	"github.com/google/uuid"
)

type AttendeeRepository interface {
	InsertAttendees(ctx context.Context, eventID uuid.UUID, userIDs []string) error
	DeleteAttendees(ctx context.Context, eventID uuid.UUID, userIDs []string) error
	UpdateAttendees(ctx context.Context, eventID uuid.UUID, userIDs []string) error
}
