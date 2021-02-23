package priceservice

import (
	"context"
	"github.com/siller174/goodsGateway/pkg/gateway/repository"
)

type PriceService struct {
	dao *dao
}

func NewPriceService(dao *dao) *PriceService {
	return &PriceService{
		dao: dao,
	}
}

func (i *PriceService) Create(ctx context.Context, shopID, itemId, price int) error {
	return i.dao.Create(ctx, shopID, itemId, price)
}

func (i *PriceService) GetLastPrice(ctx context.Context, shopId, itemId int) (int, error) {
	return i.dao.GetLastPrice(ctx, shopId, itemId)
}

func (i *PriceService) GetDAO() *repository.Mysql {
	return i.dao.repos
}
