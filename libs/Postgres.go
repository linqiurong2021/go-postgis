package libs

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/linqiurong2021/go-postgis/conf"
)

// Postgre Postgre
type Postgre struct{}

// NewPostgre NewPostgre
func NewPostgre(config *conf.Config) *Postgre {
	return &Postgre{}
}

// Connect 连接
func (p *Postgre) Connect() (dbpool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s pool_min_conns=%d pool_max_conns=%d pool_max_conn_lifetime=%d",
		conf.Conf.DB.User, conf.Conf.DB.Password, conf.Conf.DB.Host, conf.Conf.DB.Port, conf.Conf.DB.Database, conf.Conf.DB.MinPoolSize, conf.Conf.DB.MaxPoolSize, conf.Conf.DB.LifeTime)
	fmt.Printf("dsn: %s\n", dsn)

	dbpool, err = pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	defer dbpool.Close()

	return dbpool, nil
}
