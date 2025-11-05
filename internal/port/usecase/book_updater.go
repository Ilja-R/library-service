package usecase

import (
	"context"

	"github.com/Ilja-R/library-service/internal/domain"
)
type BookUpdater interface{
	UpdateBook(ctx context.Context,book domain.UpdateBookBody,id int)error
}