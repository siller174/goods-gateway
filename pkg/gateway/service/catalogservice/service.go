package catalogservice

import "context"

type Service struct {
	dao *dao
}

func NewService(dao *dao) *Service {
	return &Service{
		dao: dao,
	}
}

func (c *Service) Get(ctx context.Context, shopId int) ([]Item, error) {
	return c.dao.Get(ctx, shopId)
}
