package repository

import (
	"time"

	"gorm.io/gorm"
)

type LotteryRepositoryImpl struct {
	db *gorm.DB
}

func NewLotteryRepositoryImpl(db *gorm.DB) *LotteryRepositoryImpl {
	return &LotteryRepositoryImpl{
		db: db,
	}
}

type Lottery struct {
	LotteryID string    `gorm:"column:lottery_id;type:char(36);primaryKey;not null" json:"lottery_id"`
	EventID   string    `gorm:"column:event_id;type:char(36);not null" json:"event_id"`
	Title     string    `gorm:"column:title;type:varchar(100);not null" json:"title"`
	IsDeleted bool      `gorm:"column:is_deleted;type:boolean;not null" json:"is_deleted"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`

	// Associations
	Event   Event    `gorm:"foreignKey:EventID;references:EventID" json:"event,omitempty"`
	Winners []Winner `gorm:"foreignKey:LotteryID;references:LotteryID" json:"winners,omitempty"`
}
