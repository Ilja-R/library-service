package usecase

import (
	"github.com/Ilja-R/library-service/internal/adapter/driven/cache"
	"github.com/Ilja-R/library-service/internal/adapter/driven/dbstore"
	"github.com/Ilja-R/library-service/internal/config"
	"github.com/Ilja-R/library-service/internal/port/usecase"
	bookgetter "github.com/Ilja-R/library-service/internal/usecase/book_getter"
)
type UseCases struct{
	BookGetter usecase.BookGetter
}


func New(cfg *config.Config, store *dbstore.DBStore, cache *cache.Cache) *UseCases {
	return &UseCases{
		BookGetter: bookgetter.New(cfg,store,cache),
	}
}