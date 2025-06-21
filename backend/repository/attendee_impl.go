package repository

import (
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
