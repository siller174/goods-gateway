package articleservice

import "context"

type Service struct {
	dao *dao
}

func NewService(dao *dao) *Service {
	return &Service{
		dao: dao,
	}
}

func (i *Service) IsExist(ctx context.Context, article string, shopId int) (bool, error) {
	return i.dao.IsExist(ctx, article, shopId)
}
