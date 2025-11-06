package booksearcher

import (
	"context"

	"github.com/Ilja-R/library-service/internal/config"
	"github.com/Ilja-R/library-service/internal/domain"
	"github.com/Ilja-R/library-service/internal/port/driven"
)

type UseCase struct {
	cfg         *config.Config
	BookStorage driven.BookStorage
}

func New(cfg *config.Config, b driven.BookStorage) *UseCase {
	return &UseCase{
		cfg:         cfg,
		BookStorage: b,
	}
}

func (u *UseCase) SearchByTitle(ctx context.Context, title string) (books []domain.Book, err error) {
	books, err = u.BookStorage.SearchByTitle(ctx,title)
	if err != nil {
		return nil, err
	}
	return books, nil
}
