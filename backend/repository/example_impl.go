package repository

import "gorm.io/gorm"

type ExampleRepositoryImpl struct {
	db *gorm.DB
}

func NewExampleRepositoryImpl(db *gorm.DB) *ExampleRepositoryImpl {
	return &ExampleRepositoryImpl{
		db: db,
	}
}

func (e *ExampleRepositoryImpl) ExampleRepositoryMethod() error {
	return nil
}
