package service

import (
	"context"

	"github.com/google/uuid"
)

type EventService interface {
	CreateEvent(ctx context.Context, newEvent EventOnCreate) (uuid.UUID, error)
	GetEvents(ctx context.Context, ifDeleted bool, userID string) ([]EventSummary, error)
	GetEvent(ctx context.Context, eventID uuid.UUID, userID string) (EventDetail, error)
	EditEvent(ctx context.Context, eventID uuid.UUID, event EventOnEdit) error
	DeleteEvent(ctx context.Context, eventID uuid.UUID) error
}
