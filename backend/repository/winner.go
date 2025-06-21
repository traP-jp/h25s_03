package repository

import (
	"context"
)

type WinnerRepository interface {
	InsertWinner(ctx context.Context, winner Winner) error
}
