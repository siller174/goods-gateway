package priceservice

import (
	"context"
	"database/sql"
	"github.com/siller174/goodsGateway/pkg/gateway/api/http/errors"
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

func (dao *dao) Create(ctx context.Context, shopID, itemId, price int) error {

	q := `INSERT INTO price (ShopID, ItemID, Price) VALUES (?, ?, ?)`

	_, err := dao.repos.Db.ExecContext(ctx, q, shopID, itemId, price)

	if err != nil {
		return err
	}

	return nil
}

func (dao *dao) GetLastPrice(ctx context.Context, shopId, itemId int) (int, error) {
	var data int

	q := `SELECT Price FROM price where ItemID = ? and ShopID = ? ORDER BY PriceDate desc limit 1`

	err := dao.repos.Db.GetContext(ctx, &data, q, itemId, shopId)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.NewNotFound(err)
		}
		return 0, err
	}

	return data, nil
}
