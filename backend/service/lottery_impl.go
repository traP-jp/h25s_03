package service

import (
	"github.com/eraxyso/go-template/api"
	"github.com/eraxyso/go-template/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type LotteryServiceImpl struct {
	lotteryRepository repository.LotteryRepository
}

func NewLotteryServiceImpl(lotteryRepository repository.LotteryRepository) *LotteryServiceImpl {
	return &LotteryServiceImpl{}
}

func (ls *LotteryServiceImpl) GetLotteries(ctx echo.Context, eventID uuid.UUID) ([]api.Lottery, error) {
	lotteries, err := ls.lotteryRepository.GetLotteries(ctx, eventID)
	if err != nil {
		return nil, err
	}
	return lotteries, nil
}
