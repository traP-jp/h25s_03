package repository

import (
	"gorm.io/gorm"
)

type WinnerRepositoryImpl struct {
	db *gorm.DB
}

func NewWinnerRepositoryImpl(db *gorm.DB) *WinnerRepositoryImpl {
	return &WinnerRepositoryImpl{
		db: db,
	}
}

type Winner struct {
	LotteryID string `gorm:"column:lottery_id;type:char(36);primaryKey;not null" json:"lottery_id"`
	TraqID    string `gorm:"column:traq_id;type:varchar(32);primaryKey;not null" json:"traq_id"`
	EventID   string `gorm:"column:event_id;type:char(36);not null" json:"event_id"`

	// Associations
	Lottery Lottery `gorm:"foreignKey:LotteryID;references:LotteryID" json:"lottery,omitempty"`
	Event   Event   `gorm:"foreignKey:EventID;references:EventID" json:"event,omitempty"`
}
