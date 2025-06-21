package repository

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AttendeeRepositoryImpl struct {
	db *gorm.DB
}

func NewAttendeeRepositoryImpl(db *gorm.DB) *AttendeeRepositoryImpl {
	return &AttendeeRepositoryImpl{
		db: db,
	}
}

type Attendee struct {
	EventID string `gorm:"column:event_id;type:char(36);primaryKey;not null" json:"event_id"`
	TraqID  string `gorm:"column:traq_id;type:varchar(32);primaryKey;not null" json:"traq_id"`

	// Association
	Event Event `gorm:"foreignKey:EventID;references:EventID" json:"event,omitempty"`
}

func (ar *AttendeeRepositoryImpl) InsertAttendees(ctx echo.Context, eventID uuid.UUID, userIDs []string) error {
	attendees := make([]Attendee, 0, len(userIDs))
	for _, userID := range userIDs {
		attendees = append(attendees, Attendee{
			EventID: eventID.String(),
			TraqID:  userID,
		})
	}
	if err := ar.db.WithContext(ctx.Request().Context()).Create(&attendees).Error; err != nil {
		return err
	}
	return nil
}

func (ar *AttendeeRepositoryImpl) DeleteAttendees(ctx echo.Context, eventID uuid.UUID, userIDs []string) error {
	if err := ar.db.WithContext(ctx.Request().Context()).Where("event_id = ? AND traq_id IN (?)", eventID.String(), userIDs).Delete(&Attendee{}).Error; err != nil {
		return err
	}
	return nil
}
