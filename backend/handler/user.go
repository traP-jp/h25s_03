package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetMeJSONResponseBody struct {
	Name string `json:"name"`
}

func (h *Handler) GetMe(ctx echo.Context) error {
	userID := h.MiddlewareService.GetUserID(ctx)

	return ctx.JSON(http.StatusOK, GetMeJSONResponseBody{
		Name: userID,
	})
}
