package service

import "github.com/labstack/echo/v4"

type MiddlewareService interface {
	GetUserID(echo.Context) string
}
