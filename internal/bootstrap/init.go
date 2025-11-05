package bootstrap

import (
	"fmt"

	"net/http"

	http2 "github.com/Ilja-R/library-service/internal/adapter/driving/http"
	"github.com/Ilja-R/library-service/internal/config"
	"github.com/Ilja-R/library-service/internal/usecase"
	"github.com/redis/go-redis/v9"

	_ "github.com/lib/pq"
	
	"github.com/jmoiron/sqlx"
)

func initDB(cfg config.Postgres) (*sqlx.DB, error) {
	postgresOpen := fmt.Sprintf(
		`host=%s
			user=%s
			password=%s
			dbname=%s
			sslmode=disable`,
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)
	db, err := sqlx.Open("postgres", postgresOpen)
	if err != nil {
		return nil,err
	}

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