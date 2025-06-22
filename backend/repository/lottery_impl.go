package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LotteryRepositoryImpl struct {
	db *gorm.DB
}

var _ LotteryRepository = &LotteryRepositoryImpl{}

func NewLotteryRepositoryImpl(db *gorm.DB) *LotteryRepositoryImpl {
	return &LotteryRepositoryImpl{
		db: db,
	}
}

type lotteryModel struct {
	LotteryID string    `gorm:"column:lottery_id;type:char(36);primaryKey;not null"`
	EventID   string    `gorm:"column:event_id;type:char(36);not null"`
	Title     string    `gorm:"column:title;type:varchar(100);not null"`
	IsDeleted bool      `gorm:"column:is_deleted;type:boolean;not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;autoUpdateTime"`

	// Associations
	Winners []winnerModel `gorm:"foreignKey:LotteryID;references:LotteryID"`
}

func (lotteryModel) TableName() string {
	return "lotteries"
}

type LotteryOnCreate struct {
	Title string `json:"title"`
}

func (ls *LotteryRepositoryImpl) InsertLottery(ctx context.Context, eventID uuid.UUID, lottery LotteryOnCreate) (uuid.UUID, error) {
	lotteryID := uuid.New()
	newLottery := &lotteryModel{
		LotteryID: lotteryID.String(),
		EventID:   eventID.String(),
		Title:     lottery.Title,
		IsDeleted: false,
	}

	if err := ls.db.WithContext(ctx).Create(newLottery).Error; err != nil {
		return uuid.Nil, fmt.Errorf("insert lottery (repository): %w", err)
	}

	return lotteryID, nil
}

type LotteryWithWinners struct {
	LotteryID uuid.UUID
	EventID   uuid.UUID
	Title     string
	IsDeleted bool
	CreatedAt time.Time
	UpdatedAt time.Time
	Winners   []string
}

func (lr *LotteryRepositoryImpl) GetLotteries(ctx context.Context, eventID uuid.UUID, ifDeleted bool) ([]LotteryWithWinners, error) {
	var lotteries []lotteryModel
	query := lr.db.WithContext(ctx).Preload("Winners").Where("event_id = ?", eventID)
	if !ifDeleted {
		query = query.Where("is_deleted = ?", false)
	}
	if err := query.Find(&lotteries).Error; err != nil {
		return nil, fmt.Errorf("get lotteries (repository): %w", err)
	}

	var lotteriesResult []LotteryWithWinners
	for _, lottery := range lotteries {
		lotteryUUID, err := uuid.Parse(lottery.LotteryID)
		if err != nil {
			return nil, fmt.Errorf("parse lottery id (repository): %w", err)
		}
		eventUUID, err := uuid.Parse(lottery.EventID)
		if err != nil {
			return nil, fmt.Errorf("parse event id (repository): %w", err)
		}
		l := LotteryWithWinners{
			LotteryID: lotteryUUID,
			EventID:   eventUUID,
			Title:     lottery.Title,
			IsDeleted: lottery.IsDeleted,
			CreatedAt: lottery.CreatedAt,
			UpdatedAt: lottery.UpdatedAt,
		}
		l.Winners = make([]string, len(lottery.Winners))
		for i, winner := range lottery.Winners {
			l.Winners[i] = winner.TraqID
		}
		lotteriesResult = append(lotteriesResult, l)
	}

	return lotteriesResult, nil
}

func (ls *LotteryRepositoryImpl) GetLottery(ctx context.Context, eventID uuid.UUID, lotteryID uuid.UUID) (LotteryWithWinners, error) {
	var lottery lotteryModel
	err := ls.db.WithContext(ctx).Preload("Winners").Where("lottery_id = ? AND event_id = ?", lotteryID, eventID).Where("is_deleted = ?", false).First(&lottery).Error
	if err != nil {
		return LotteryWithWinners{}, fmt.Errorf("get lottery (repository): %w", err)
	}

	lotteryUUID, err := uuid.Parse(lottery.LotteryID)
	if err != nil {
		return LotteryWithWinners{}, fmt.Errorf("parse lottery id (repository): %w", err)
	}
	eventUUID, err := uuid.Parse(lottery.EventID)
	if err != nil {
		return LotteryWithWinners{}, fmt.Errorf("parse event id (repository): %w", err)
	}

	foundLottery := LotteryWithWinners{
		LotteryID: lotteryUUID,
		EventID:   eventUUID,
		Title:     lottery.Title,
		IsDeleted: lottery.IsDeleted,
		CreatedAt: lottery.CreatedAt,
		UpdatedAt: lottery.UpdatedAt,
	}
	foundLottery.Winners = make([]string, len(lottery.Winners))
	for i, winner := range lottery.Winners {
		foundLottery.Winners[i] = winner.TraqID
	}

	return foundLottery, nil
}

func (lr *LotteryRepositoryImpl) DeleteLottery(ctx context.Context, lotteryID uuid.UUID) error {
	if err := lr.db.WithContext(ctx).Model(&lotteryModel{}).Where("lottery_id = ?", lotteryID).Update("is_deleted", true).Error; err != nil {
		return fmt.Errorf("delete lottery (repository): %w", err)
	}
	return nil
}
