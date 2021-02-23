package shopservice

import (
	"context"
	"github.com/siller174/goodsGateway/pkg/gateway/repository"
)

type dao struct {
	repos *repository.Mysql
}

func NewDAO(mysql *repository.Mysql) *dao {
	return &dao{
		repos: mysql,
	}
}

func (dao *dao) GetShopByID(ctx context.Context, shopID int) (*string, error) {
	var res = new(string)
	q := `SELECT ShopName FROM shop WHERE ShopID = ?`

	err := dao.repos.Db.GetContext(ctx, &res, q, shopID)

	if err != nil {
		return nil, err
	}
	return res, err
}
