package catalogservice

import (
	"context"
	"database/sql"
	"fmt"
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

func (dao *dao) Get(ctx context.Context, i int) ([]Item, error) {
	var data []Item

	q := `SELECT Url, Count, CategoryID FROM catalogs where Enabled = 1 and ShopID = ?`

	err := dao.repos.Db.SelectContext(ctx, &data, q, i)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFound(err)
		}
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.NewNotFound(fmt.Errorf("catalog with id %v not found", i))
	}

	return data, nil
}
