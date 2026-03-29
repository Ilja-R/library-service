package bootstrap

import (
	"context"
	"fmt"

	"github.com/Ilja-R/library-service/internal/adapter/driven/cache"
	"github.com/Ilja-R/library-service/internal/adapter/driven/dbstore"
	"github.com/Ilja-R/library-service/internal/config"
	"github.com/Ilja-R/library-service/internal/usecase"
)

func initLayers(cfg config.Config) *App {
	teardown := make([]func(), 0)
	db, err := initDB(*cfg.Postgres)
	if err != nil {
		panic(err)
	}
	storage := dbstore.New(db)
	rdb := initRedis(cfg.Redis)
	cacheStorage := cache.NewCache(rdb)
	teardown = append(teardown, func() {
		if err := db.Close(); err != nil {
			fmt.Println(err)
		}
	})
	uc := usecase.New(&cfg, storage, cacheStorage)
	httpSrv := initHTTPService(&cfg, uc)
	teardown = append(teardown,
		func() {

			ctxShutDown, cancel := context.WithTimeout(context.Background(), gracefulDeadline)
			defer cancel()
			if err := httpSrv.Shutdown(ctxShutDown); err != nil {
				return
			}

		},
	)

	return &App{
		cfg:      cfg,
		rest:     httpSrv,
		teardown: teardown,
	}
}