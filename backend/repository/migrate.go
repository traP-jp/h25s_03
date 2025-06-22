package repository

import (
	"fmt"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&eventModel{},
		&adminModel{},
		&attendeeModel{},
		&lotteryModel{},
		&winnerModel{},
	)
	if err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}

	return nil
}
