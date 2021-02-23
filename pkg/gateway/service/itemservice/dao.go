package itemservice

import (
	"context"
	"database/sql"
	"github.com/siller174/goodsGateway/pkg/gateway/api/http/errors"
	"github.com/siller174/goodsGateway/pkg/gateway/repository"
	"github.com/siller174/goodsGateway/pkg/gateway/structs"
)

type dao struct {
	repos *repository.Mysql
}

func NewDAO(mysql *repository.Mysql) *dao {
	return &dao{
		repos: mysql,
	}
}

func (dao *dao) GetId(ctx context.Context, itemName string, categoryID int) (int, error) {
	var data int

	q := `SELECT ItemID FROM item WHERE ItemName = ? AND CategoryID = ?`

	err := dao.repos.Db.GetContext(ctx, &data, q, itemName, categoryID)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.NewNotFound(err)
		}
		return 0, err
	}

	return data, nil
}

func (dao *dao) Create(ctx context.Context, good structs.Good) error {
	q := `CALL createNewGood(?, ?, ?, ?, ?, ?, ?)`

	_, err := dao.repos.Db.ExecContext(ctx, q, good.ShopId, good.CategoryID, good.Name, good.Article, good.Price, good.URL, good.ImageURL)

	if err != nil {
		return errors.NewInternalErr(err)
	}
	return nil
}

func (dao *dao) IsExist(ctx context.Context, name string) (bool, error) {
	var i int
	q := "SELECT count(*) FROM item WHERE ItemName = ?"

	err := dao.repos.Db.GetContext(ctx, &i, q, name)

	if err != nil {
		return false, err
	}

	return i != 0, nil
}
