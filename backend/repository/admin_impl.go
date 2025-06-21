package repository

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AdminRepositoryImpl struct {
	db *gorm.DB
}

func NewAdminRepositoryImpl(db *gorm.DB) *AdminRepositoryImpl {
	return &AdminRepositoryImpl{
		db: db,
	}
}

type Admin struct {
	EventID string `gorm:"column:event_id;type:char(36);primaryKey;not null" json:"event_id"`
	TraqID  string `gorm:"column:traq_id;type:varchar(32);primaryKey;not null" json:"traq_id"`

	// Association
	Event Event `gorm:"foreignKey:EventID;references:EventID" json:"event,omitempty"`
}

func (ar *AdminRepositoryImpl) InsertAdmins(ctx echo.Context, eventID uuid.UUID, userIDs []string) error {
	admins := make([]Admin, 0, len(userIDs))
	for _, userID := range userIDs {
		admins = append(admins, Admin{
			EventID: eventID.String(),
			TraqID:  userID,
		})
	}
	if err := ar.db.WithContext(ctx.Request().Context()).Create(&admins).Error; err != nil {
		return err
	}
	return nil
}

func (ar *AdminRepositoryImpl) DeleteAdmins(ctx echo.Context, eventID uuid.UUID, userIDs []string) error {
	if err := ar.db.WithContext(ctx.Request().Context()).Where("event_id = ? AND traq_id IN (?)", eventID.String(), userIDs).Delete(&Admin{}).Error; err != nil {
		return err
	}
	return nil
}
