package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WinnerRepositoryImpl struct {
	db *gorm.DB
}

var _ WinnerRepository = &WinnerRepositoryImpl{}

func NewWinnerRepositoryImpl(db *gorm.DB) *WinnerRepositoryImpl {
	return &WinnerRepositoryImpl{
		db: db,
	}
}

type winnerModel struct {
	LotteryID string `gorm:"column:lottery_id;type:char(36);primaryKey;not null"`
	TraqID    string `gorm:"column:traq_id;type:varchar(32);primaryKey;not null"`
	EventID   string `gorm:"column:event_id;type:char(36);not null"`
}

func (winnerModel) TableName() string {
	return "winners"
}

type Winner struct {
	EventID   uuid.UUID
	LotteryID uuid.UUID
	TraqID    string
}

func (wr *WinnerRepositoryImpl) InsertWinner(ctx context.Context, winner Winner) error {
	newWinner := winnerModel{
		EventID:   winner.EventID.String(),
		LotteryID: winner.LotteryID.String(),
		TraqID:    winner.TraqID,
	}
	if err := wr.db.WithContext(ctx).Create(&newWinner).Error; err != nil {
		return fmt.Errorf("insert winner (repository): %w", err)
	}
	return nil
}
