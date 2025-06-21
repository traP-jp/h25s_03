package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/eraxyso/go-template/api"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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

func (ls *LotteryRepositoryImpl) InsertLottery(ctx context.Context, eventID uuid.UUID, lotteryBody api.PostEventJSONRequestBody) (uuid.UUID, error) {

	newLottery := &Lottery{
		LotteryID: uuid.NewString(),

		EventID: eventID.String(),
		Title:   lotteryBody.Title,
	}

	if err := ls.db.WithContext(ctx).Create(newLottery).Error; err != nil {
		return uuid.Nil, fmt.Errorf("failed to create lottery in db: %w", err)
	}

	createdID, err := uuid.Parse(newLottery.LotteryID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to parse created lottery id: %w", err)
	}

	return createdID, nil
}

func (lr *LotteryRepositoryImpl) GetLotteries(ctx echo.Context, eventID uuid.UUID, ifDeleted bool) ([]api.Lottery, error) {
	var lotteries []Lottery
	query := lr.db.WithContext(ctx.Request().Context()).Preload("Winners").Where("event_id = ?", eventID)
	if !ifDeleted {
		query = query.Where("is_deleted = ?", false)
	}
	if err := query.Find(&lotteries).Error; err != nil {
		return nil, err
	}

	var apiLotteries []api.Lottery
	for _, lottery := range lotteries {
		lotteryUUID, err := uuid.Parse(lottery.LotteryID)
		if err != nil {
			return nil, err
		}
		eventUUID, err := uuid.Parse(lottery.EventID)
		if err != nil {
			return nil, err
		}
		apiLotteries = append(apiLotteries, api.Lottery{
			LotteryId: lotteryUUID,
			EventId:   eventUUID,
			Title:     lottery.Title,
			IsDeleted: lottery.IsDeleted,
			CreatedAt: lottery.CreatedAt,
			UpdatedAt: lottery.UpdatedAt,
		})
	}

	return apiLotteries, nil
}

func (lr *LotteryRepositoryImpl) DeleteLottery(ctx echo.Context, lotteryID uuid.UUID) error {
	if err := lr.db.WithContext(ctx.Request().Context()).Model(&Lottery{}).Where("lottery_id = ?", lotteryID).Update("is_deleted", true).Error; err != nil {
		return err
	}
	return nil
}
