package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/linqiurong2021/go-postgis/conf"
)

// PostgreDB PostgreDB
type PostgreDB struct {
	Conf *conf.Config
}

// NewPostgreDB NewPostgreDB
func NewPostgreDB(config *conf.Config) *PostgreDB {
	return &PostgreDB{
		Conf: config,
	}
}

// Connect 连接
func (p *PostgreDB) Connect() (dbpool *pgxpool.Pool, err error) {
	//
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s pool_min_conns=%d pool_max_conns=%d pool_max_conn_lifetime=%d",
		p.Conf.DB.User, p.Conf.DB.Password, p.Conf.DB.Host, p.Conf.DB.Port, p.Conf.DB.Database, p.Conf.DB.MinPoolSize, p.Conf.DB.MaxPoolSize, p.Conf.DB.LifeTime)
	fmt.Printf("dsn %s\n", dsn)
	dbpool, err = pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return dbpool, nil
}
