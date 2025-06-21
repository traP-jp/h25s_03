package service

import "github.com/labstack/echo/v4"

type MiddlewareService interface {
	MiddlewareServiceExample(next echo.HandlerFunc) echo.HandlerFunc
}
