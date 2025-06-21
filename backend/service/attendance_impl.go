package service

import (
	"github.com/eraxyso/go-template/repository"
)

type AttendanceServiceImpl struct {
	eventRepository repository.EventRepository
}

func NewAttendanceServiceImpl(exampleRepository repository.EventRepository) *AttendanceServiceImpl {
	return &AttendanceServiceImpl{
		eventRepository: exampleRepository,
	}
}
