package driven

import (
	"context"

	"github.com/Ilja-R/library-service/internal/domain"
)
 





type BookStorage interface {
	GetAllBooks(ctx context.Context) (books []domain.Book, err error)
	GetBookByID(ctx context.Context, id int) (domain.Book, error)
	CreateProduct(ctx context.Context, book domain.Book) (err error)
	UpdateProductByID(ctx context.Context, book domain.Book) (err error)
	DeleteBookByID(ctx context.Context, id int) (err error)
}