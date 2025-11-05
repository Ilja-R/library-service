package dbstore

import (
	"context"


	"github.com/Ilja-R/library-service/internal/domain"
	"github.com/jmoiron/sqlx"
)

type DBStore struct {
	BookStorage *BookStorage
}

// CreateProduct implements driven.BookStorage.
func (d *DBStore) CreateBook(ctx context.Context, createBook domain.CreateBook) (error) {
	err:=d.BookStorage.CreateProduct(ctx,createBook)
	if err!=nil{
		return err 
	}
	return nil 
}

// DeleteBookByID implements driven.BookStorage.
func (d *DBStore) DeleteBookByID(ctx context.Context, id int) error {
	err:=d.BookStorage.DeleteBookByID(ctx,id)
	if err!=nil{
		return err 
	}
	return nil
}

// GetAllBooks implements driven.BookStorage.
func (d *DBStore) GetAllBooks(ctx context.Context) (books []domain.Book, err error) {
	books,err=d.BookStorage.GetAllBooks(ctx)
	
	if err!=nil{
		return nil ,err
	}
	return 
}

// GetBookByID implements driven.BookStorage.
func (d *DBStore) GetBookByID(ctx context.Context, id int) (domain.Book, error) {
	book,err:=d.BookStorage.GetBookByID(ctx,id)
	if err!=nil{
		return domain.Book{} , err
	}
	return book ,nil
}

// UpdateProductByID implements driven.BookStorage.
func (d *DBStore) UpdateBookByID(ctx context.Context, book domain.UpdateBookBody,id int ) error {
	err:=d.BookStorage.UpdateBookByID(ctx,book ,id)
	if err!=nil{
		return err 
	}
	return nil 

}

func New(db *sqlx.DB) *DBStore {
	return &DBStore{
		BookStorage: NewBookStorage(db),
	}
}
