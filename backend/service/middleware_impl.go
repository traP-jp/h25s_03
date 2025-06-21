package service

import "github.com/labstack/echo/v4"

type MiddlewareServiceImpl struct {
}

func NewMiddlewareServiceImpl() *MiddlewareServiceImpl {
	return &MiddlewareServiceImpl{}
}

func GetUserID(ctx echo.Context) string {
	userID := ctx.Request().Header.Get("X-Forwarded-User")
	if userID == "" {
		userID = "cp20"
	}
	return userID
}
