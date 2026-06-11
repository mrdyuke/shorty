package pool

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mrdyuke/shorty/config"
)

func NewPostgresPool(cfg *config.Config) (*pgxpool.Pool, error) {

	poolCfg, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, err
	}

	poolCfg.ConnConfig.Port = cfg.DBPort
	poolCfg.ConnConfig.Host = cfg.DBHost
	poolCfg.ConnConfig.User = cfg.DBUser
	poolCfg.ConnConfig.Password = cfg.DBPassword
	poolCfg.ConnConfig.Database = cfg.DBName

	poolCfg.MaxConns = cfg.MaxConnections
	poolCfg.MaxConnLifetime = cfg.MaxLifetime

	pool, err := pgxpool.NewWithConfig(context.Background(), poolCfg)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return pool, nil

}
