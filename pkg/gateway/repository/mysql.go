package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/siller174/goodsGateway/pkg/gateway/config"
)

type Mysql struct {
	Db *sqlx.DB
}

func NewMysql(ctx context.Context, cfg config.Db) (*Mysql, error) {
	sourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.UserName, cfg.Password, cfg.Host, cfg.DbName)
	db, err := sqlx.ConnectContext(ctx, "mysql", sourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Mysql{
		Db: db,
	}, nil
}
