package usecase

import (
	"context"

	"github.com/Ilja-R/library-service/internal/domain"
)
type BookGetter interface {
	GetAllBooks(ctx context.Context)([]domain.Book,error)
	GetBookByID(ctx context.Context,id int )(book domain.Book,err error)
}