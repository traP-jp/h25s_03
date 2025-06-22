package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminRepositoryImpl struct {
	db *gorm.DB
}

var _ AdminRepository = &AdminRepositoryImpl{}

func NewAdminRepositoryImpl(db *gorm.DB) *AdminRepositoryImpl {
	return &AdminRepositoryImpl{
		db: db,
	}
}

type adminModel struct {
	EventID string `gorm:"column:event_id;type:char(36);primaryKey;not null"`
	TraqID  string `gorm:"column:traq_id;type:varchar(32);primaryKey;not null"`
}

func (adminModel) TableName() string {
	return "admins"
}

func (ar *AdminRepositoryImpl) InsertAdmins(ctx context.Context, eventID uuid.UUID, userIDs []string) error {
	admins := make([]adminModel, 0, len(userIDs))
	for _, userID := range userIDs {
		admins = append(admins, adminModel{
			EventID: eventID.String(),
			TraqID:  userID,
		})
	}
	if err := ar.db.WithContext(ctx).Create(&admins).Error; err != nil {
		return fmt.Errorf("insert admins (repository): %w", err)
	}
	return nil
}

func (ar *AdminRepositoryImpl) DeleteAdmins(ctx context.Context, eventID uuid.UUID, userIDs []string) error {
	if err := ar.db.WithContext(ctx).Where("event_id = ? AND traq_id IN ?", eventID.String(), userIDs).Delete(&adminModel{}).Error; err != nil {
		return fmt.Errorf("delete admins (repository): %w", err)
	}
	return nil
}

func (ar *AdminRepositoryImpl) UpdateAdmins(ctx context.Context, eventID uuid.UUID, userIDs []string) error {
	if err := ar.db.WithContext(ctx).Where("event_id = ?", eventID.String()).Delete(&adminModel{}).Error; err != nil {
		return fmt.Errorf("delete admins (repository): %w", err)
	}
	if err := ar.InsertAdmins(ctx, eventID, userIDs); err != nil {
		return fmt.Errorf("insert admins (repository): %w", err)
	}
	return nil
}

func (ar *AdminRepositoryImpl) CheckAdmin(ctx context.Context, eventID uuid.UUID, userID string) (bool, error) {
	var count int64
	if err := ar.db.WithContext(ctx).Model(&adminModel{}).
		Where("event_id = ? AND traq_id = ?", eventID.String(), userID).
		Count(&count).Error; err != nil {
		return false, fmt.Errorf("check admin (repository): %w", err)
	}
	return count > 0, nil
}
