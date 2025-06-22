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
	LotteryID string `gorm:"column:lottery_id;type:char(36);primaryKey;not null" json:"lottery_id"`
	TraqID    string `gorm:"column:traq_id;type:varchar(32);primaryKey;not null" json:"traq_id"`
	EventID   string `gorm:"column:event_id;type:char(36);not null" json:"event_id"`

	// Associations
	Lottery lotteryModel `gorm:"foreignKey:LotteryID;references:LotteryID" json:"lottery,omitempty"`
	Event   eventModel   `gorm:"foreignKey:EventID;references:EventID" json:"event,omitempty"`
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

func (wr *WinnerRepositoryImpl) GetEventWinners(ctx context.Context, eventID uuid.UUID) ([]string, error) {
	var winners []string
	err := wr.db.WithContext(ctx).
		Model(&winnerModel{}).
		Where("event_id = ?", eventID).
		Pluck("traq_id", &winners).Error
	if err != nil {
		return nil, fmt.Errorf("get event winners (repository): %w", err)
	}
	return winners, nil
}

func (wr *WinnerRepositoryImpl) GetLotteryWinnners(ctx context.Context, lotteryID uuid.UUID) ([]string, error) {
	var winners []string
	err := wr.db.WithContext(ctx).
		Model(&winnerModel{}).
		Where("lottery_id = ?", lotteryID).
		Pluck("traq_id", &winners).Error
	if err != nil {
		return nil, fmt.Errorf("get lottery winners: %w", err)
	}
	return winners, nil
}
