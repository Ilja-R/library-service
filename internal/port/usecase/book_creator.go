package usecase

import (
	"context"

	"github.com/Ilja-R/library-service/internal/domain"
)
type BookCreator interface {
	CreateBook(ctx context.Context,createBook domain.CreateBook)error
}