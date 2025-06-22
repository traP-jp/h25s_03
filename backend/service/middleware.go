package service

import "github.com/labstack/echo/v4"

type MiddlewareService interface {
	GetUserID(echo.Context) string
	EventAdminAuthentication(next echo.HandlerFunc) echo.HandlerFunc
	EventRegistrationAuthentication(next echo.HandlerFunc) echo.HandlerFunc
}
