package shopservice

import "context"

type Service struct {
	dao *dao
}

func NewService(dao *dao) *Service {
	return &Service{
		dao: dao,
	}
}

func (i *Service) GetShopByID(ctx context.Context, shopID int) (*string, error) {
	return i.dao.GetShopByID(ctx, shopID)
}
