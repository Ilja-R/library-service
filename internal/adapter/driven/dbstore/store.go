package dbstore

import (
	"context"
	"github.com/Ilja-R/library-service/internal/domain"
	"github.com/jmoiron/sqlx"
)

type DBStore struct {
	ProductStorage *BookStorage
}

// CreateProduct implements driven.BookStorage.
func (d *DBStore) CreateProduct(ctx context.Context, book domain.Book) (err error) {
	panic("unimplemented")
}

// DeleteBookByID implements driven.BookStorage.
func (d *DBStore) DeleteBookByID(ctx context.Context, id int) (err error) {
	panic("unimplemented")
}

// GetAllBooks implements driven.BookStorage.
func (d *DBStore) GetAllBooks(ctx context.Context) (books []domain.Book, err error) {
	panic("unimplemented")
}

// GetBookByID implements driven.BookStorage.
func (d *DBStore) GetBookByID(ctx context.Context, id int) (domain.Book, error) {
	panic("unimplemented")
}

// UpdateProductByID implements driven.BookStorage.
func (d *DBStore) UpdateProductByID(ctx context.Context, book domain.Book) (err error) {
	panic("unimplemented")
}

func New(db *sqlx.DB) *DBStore {
	return &DBStore{
		ProductStorage: NewBookStorage(db),
	}
}
