package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AttendeeRepositoryImpl struct {
	db *gorm.DB
}

var _ AttendeeRepository = &AttendeeRepositoryImpl{}

func NewAttendeeRepositoryImpl(db *gorm.DB) *AttendeeRepositoryImpl {
	return &AttendeeRepositoryImpl{
		db: db,
	}
}

type attendeeModel struct {
	EventID string `gorm:"column:event_id;type:char(36);primaryKey;not null"`
	TraqID  string `gorm:"column:traq_id;type:varchar(32);primaryKey;not null"`
}

func (attendeeModel) TableName() string {
	return "attendees"
}

func (ar *AttendeeRepositoryImpl) InsertAttendees(ctx context.Context, eventID uuid.UUID, userIDs []string) error {
	attendees := make([]attendeeModel, 0, len(userIDs))
	for _, userID := range userIDs {
		attendees = append(attendees, attendeeModel{
			EventID: eventID.String(),
			TraqID:  userID,
		})
	}
	if err := ar.db.WithContext(ctx).Create(&attendees).Error; err != nil {
		return fmt.Errorf("insert attendees (repository): %w", err)
	}
	return nil
}

func (ar *AttendeeRepositoryImpl) DeleteAttendees(ctx context.Context, eventID uuid.UUID, userIDs []string) error {
	if err := ar.db.WithContext(ctx).Where("event_id = ? AND traq_id IN ?", eventID.String(), userIDs).Delete(&attendeeModel{}).Error; err != nil {
		return fmt.Errorf("delete attendees (repository): %w", err)
	}
	return nil
}

func (ar *AttendeeRepositoryImpl) UpdateAttendees(ctx context.Context, eventID uuid.UUID, userIDs []string) error {
	if err := ar.db.WithContext(ctx).Where("event_id = ?", eventID.String()).Delete(&attendeeModel{}).Error; err != nil {
		return fmt.Errorf("delete attendees (repository): %w", err)
	}
	if err := ar.InsertAttendees(ctx, eventID, userIDs); err != nil {
		return fmt.Errorf("insert attendees (repository): %w", err)
	}
	return nil
}

func (ar *AttendeeRepositoryImpl) GetEventAttendees(ctx context.Context, eventID uuid.UUID) ([]string, error) {
	var traqIDs []string
	err := ar.db.WithContext(ctx).
		Model(&attendeeModel{}).
		Where("event_id = ?", eventID).
		Pluck("traq_id", &traqIDs).Error
	if err != nil {
		return nil, fmt.Errorf("get sttendees: %w", err)
	}
	return traqIDs, nil
}
