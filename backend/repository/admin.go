package repository

import (
	"context"

	"github.com/google/uuid"
)

type AdminRepository interface {
	InsertAdmins(ctx context.Context, eventID uuid.UUID, userIDs []string) error
	DeleteAdmins(ctx context.Context, eventID uuid.UUID, userIDs []string) error
	UpdateAdmins(ctx context.Context, eventID uuid.UUID, userIDs []string) error
}
