package service

import "github.com/labstack/echo/v4"

type MiddlewareServiceImpl struct {
}

func NewMiddlewareServiceImpl() *MiddlewareServiceImpl {
	return &MiddlewareServiceImpl{}
}

func (ms *MiddlewareServiceImpl) MiddlewareServiceExample(next echo.HandlerFunc) echo.HandlerFunc {
	return next
}
