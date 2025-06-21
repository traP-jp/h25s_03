package service

import (
	"github.com/labstack/echo/v4"

	"github.com/eraxyso/go-template/repository"
)

type ExampleServiceImpl struct {
	exampleRepository repository.ExampleRepository
}

func NewExampleServiceImpl(exampleRepository repository.ExampleRepository) *ExampleServiceImpl {
	return &ExampleServiceImpl{
		exampleRepository: exampleRepository,
	}
}

func (e *ExampleServiceImpl) ExampleServiceMethod(c echo.Context, paramExample string, requestExample *string) error {
	return nil
}
