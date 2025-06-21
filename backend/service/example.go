package service

import "github.com/labstack/echo/v4"

type ExampleService interface {
	ExampleServiceMethod(c echo.Context, paramExample string, requestExample *string) error
}
