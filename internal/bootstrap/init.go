package bootstrap

import (
	"fmt"
	"net/http"
	http2 "github.com/Ilja-R/library-service/internal/adapter/driving/http"
	"github.com/Ilja-R/library-service/internal/config"
	"github.com/Ilja-R/library-service/internal/usecase"
	"github.com/redis/go-redis/v9"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func initDB(cfg config.Postgres) (*sqlx.DB, error) {
	connConfig, err := pgx.ParseConfig(cfg.ConnectionURL())
	if err != nil {
		return nil, err
	}

	connStr := stdlib.RegisterConnConfig(connConfig)

	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		return db, err
	}

	// Connection configuration
	// more info here https://www.alexedwards.net/blog/configuring-sqldb
	db.SetMaxOpenConns(cfg.MaxOpenConnections)
	db.SetMaxIdleConns(cfg.MaxIdleConnections)
	db.SetConnMaxLifetime(cfg.ConnectionMaxLifetime)

	return db, nil
}

func initRedis(cfg *config.Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		DB:   cfg.DB,
	})

	return rdb
}

func initHTTPService(
	cfg *config.Config,
	uc *usecase.UseCases,
) *http.Server {
	return http2.New(
		cfg,
		uc,
	)
}