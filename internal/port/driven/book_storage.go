package driven

import (
	"context"

	"github.com/Ilja-R/library-service/internal/domain"
)
 





type BookStorage interface {
	GetAllBooks(ctx context.Context) (books []domain.Book, err error)
	GetBookByID(ctx context.Context, id int) (domain.Book, error)
	CreateBook(ctx context.Context, createBook domain.CreateBook) (err error)
	UpdateBookByID(ctx context.Context, book domain.UpdateBookBody, id int ) (err error)
	DeleteBookByID(ctx context.Context, id int) (err error)
}