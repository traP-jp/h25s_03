package handler

import (
	"fmt"
	"net/http"

	"github.com/eraxyso/go-template/api"
	"github.com/eraxyso/go-template/service"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type PostLotteryJSONResponseBody struct {
	LotteryID openapi_types.UUID `json:"lottery_id"`
}

func (h *Handler) PostLottery(ctx echo.Context, eventID openapi_types.UUID) error {
	requestBody := api.PostLotteryJSONBody{}
	if err := ctx.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("bind request body (handler): %w", err))
	}
	lotteryOnCreate := service.LotteryOnCreate{
		Title: requestBody.Title,
	}
	createdID, err := h.LotteryService.CreateLottery(ctx.Request().Context(), eventID, lotteryOnCreate)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("create lottery (handler): %w", err))
	}

	return ctx.JSON(http.StatusCreated, PostLotteryJSONResponseBody{
		LotteryID: createdID,
	})
}

type GetLotteriesJSONResponseBody []api.Lottery

func (h *Handler) GetLotteries(ctx echo.Context, eventID openapi_types.UUID, params api.GetLotteriesParams) error {
	lotteries, err := h.LotteryService.GetLotteries(ctx.Request().Context(), eventID, params.IfDeleted)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("get lotteries (handler): %w", err))
	}
	response := make([]api.Lottery, len(lotteries))
	for i, lottery := range lotteries {
		response[i] = api.Lottery{
			LotteryId: lottery.LotteryID,
			EventId:   lottery.EventID,
			Title:     lottery.Title,
			IsDeleted: lottery.IsDeleted,
			CreatedAt: lottery.CreatedAt,
			UpdatedAt: lottery.UpdatedAt,
			Winners:   lottery.Winners,
		}
	}
	return ctx.JSON(http.StatusOK, GetLotteriesJSONResponseBody(response))
}

type GetLotteryJSONResponseBody api.Lottery

func (h *Handler) GetLottery(ctx echo.Context, eventID openapi_types.UUID, lotteryID openapi_types.UUID) error {
	lottery, err := h.LotteryService.GetLottery(ctx.Request().Context(), eventID, lotteryID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("get lottery (handler): %w", err))
	}
	return ctx.JSON(http.StatusOK, GetLotteryJSONResponseBody{
		LotteryId: lottery.LotteryID,
		EventId:   lottery.EventID,
		Title:     lottery.Title,
		IsDeleted: lottery.IsDeleted,
		CreatedAt: lottery.CreatedAt,
		UpdatedAt: lottery.UpdatedAt,
		Winners:   lottery.Winners,
	})
}

func (h *Handler) DeleteLottery(ctx echo.Context, eventID openapi_types.UUID, lotteryID openapi_types.UUID) error {
	err := h.LotteryService.DeleteLottery(ctx.Request().Context(), eventID, lotteryID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("delete lottery (handler): %w", err))
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (h *Handler) RollLottery(ctx echo.Context, eventID openapi_types.UUID, lotteryID openapi_types.UUID, params api.RollLotteryParams) error {
	// TODO: Implement RollLottery
	return ctx.NoContent(http.StatusNotImplemented)
}
