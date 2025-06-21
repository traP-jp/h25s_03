package handler

import (
	"net/http"

	"github.com/eraxyso/go-template/api"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handler) GetLotteries(ctx echo.Context, eventID openapi_types.UUID, params api.GetLotteriesParams) error {
	lotteries, err := h.LotteryService.GetLotteries(ctx, eventID, params.IfDeleted)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, lotteries)
}

func (h *Handler) DeleteLottery(ctx echo.Context, eventID openapi_types.UUID, lotteryID openapi_types.UUID) error {
	err := h.LotteryService.DeleteLottery(ctx, eventID, lotteryID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (h *Handler) PostLotteries(ctx echo.Context, eventID openapi_types.UUID) error {
	requestBody := api.PostLotteriesJSONRequestBody{}
	if err := ctx.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if err := h.LotteryService.CreateLottery(ctx, eventID, requestBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return ctx.NoContent(http.StatusOK)
}
