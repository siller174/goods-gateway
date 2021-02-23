package articleservice

import (
	"context"
	"github.com/siller174/goodsGateway/pkg/gateway/repository"
	"strconv"
	"strings"
)

type dao struct {
	repos *repository.Mysql
}

func NewDAO(mysql *repository.Mysql) *dao {
	return &dao{
		repos: mysql,
	}
}

func (dao *dao) IsExist(ctx context.Context, article string, shopId int) (bool, error) {
	var i int
	q := "SELECT count(*) FROM article WHERE Article = '$1' AND ShopID = $2"

	q = strings.ReplaceAll(q, "$1", article)
	q = strings.ReplaceAll(q, "$2", strconv.Itoa(shopId))

	err := dao.repos.Db.GetContext(ctx, &i, q)

	if err != nil {
		return false, err
	}

	return i != 0, nil
}
