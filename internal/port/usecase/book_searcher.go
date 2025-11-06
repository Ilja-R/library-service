package usecase

import (
	"context"

	"github.com/Ilja-R/library-service/internal/domain"
)
type BookSearcher interface {
	SearchByTitle(ctx context.Context,title string)([]domain.Book ,error)
}