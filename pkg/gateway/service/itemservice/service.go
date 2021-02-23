package itemservice

import (
	"context"
	"github.com/siller174/goodsGateway/pkg/gateway/structs"
)

type Service struct {
	dao *dao
}

func NewService(dao *dao) *Service {
	return &Service{
		dao: dao,
	}
}

func (s *Service) GetId(ctx context.Context, itemName string, categoryID int) (int, error) {
	return s.dao.GetId(ctx, itemName, categoryID)
}

func (s *Service) Create(ctx context.Context, good structs.Good) error {
	return s.dao.Create(ctx, good)
}

func (s *Service) IsExist(ctx context.Context, itemName string) (bool, error) {
	return s.dao.IsExist(ctx, itemName)
}
