package repository

import (
	"context"

	"github.com/google/uuid"
)

type EventRepository interface {
	InsertEvent(ctx context.Context, event EventOnCreate) (uuid.UUID, error)
	GetEvents(ctx context.Context, includeDeleted bool) ([]EventWithAdminsAndAttendees, error)
	GetEvent(ctx context.Context, eventID uuid.UUID) (EventWithAdminsAndAttendees, error)
	UpdateEvent(ctx context.Context, eventID uuid.UUID, event EventOnUpdate) error
	DeleteEvent(ctx context.Context, eventID uuid.UUID) error
}
