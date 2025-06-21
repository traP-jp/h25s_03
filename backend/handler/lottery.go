package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handler) GetLotteries(ctx echo.Context, eventID openapi_types.UUID) error {
	lotteries, err := h.LotteryService.GetLotteries(ctx, eventID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, lotteries)
}
