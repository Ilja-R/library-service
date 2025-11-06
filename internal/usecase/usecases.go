package usecase

import (
	"github.com/Ilja-R/library-service/internal/adapter/driven/cache"
	"github.com/Ilja-R/library-service/internal/adapter/driven/dbstore"
	"github.com/Ilja-R/library-service/internal/config"
	"github.com/Ilja-R/library-service/internal/port/usecase"
	bookcreator "github.com/Ilja-R/library-service/internal/usecase/book_creator"
	bookdeleter "github.com/Ilja-R/library-service/internal/usecase/book_deleter"
	bookgetter "github.com/Ilja-R/library-service/internal/usecase/book_getter"
	booksearcher "github.com/Ilja-R/library-service/internal/usecase/book_searcher"
	bookupdater "github.com/Ilja-R/library-service/internal/usecase/book_updater"
)
type UseCases struct{
	BookGetter usecase.BookGetter
	BookCreator usecase.BookCreator
	BookUpdater usecase.BookUpdater
	BookDeleter usecase.BookDeleter
	BookSearcher usecase.BookSearcher
}


func New(cfg *config.Config, store *dbstore.DBStore, cache *cache.Cache) *UseCases {
	return &UseCases{
		BookGetter: bookgetter.New(cfg,store,cache),
		BookCreator: bookcreator.New(cfg,store),
		BookUpdater: bookupdater.New(cfg,store),
		BookDeleter: bookdeleter.New(cfg,store),
		BookSearcher: booksearcher.New(cfg,store),
	}
}