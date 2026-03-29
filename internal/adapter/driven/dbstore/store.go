package dbstore

import (
	"context"
	"log"

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
	log.Println("store get all books",books)
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
func (d *DBStore) SearchByTitle(ctx context.Context,title string) (books []domain.Book,err error) {
	books ,err=d.BookStorage.SearchByTitle(ctx,title)
	if err!=nil{
		return nil ,err
	}
	return books ,nil  

}



func(d *DBStore)OrderBookByTitle(ctx context.Context,title string,username string)error{
	err:=d.BookStorage.OrderBookByTitle(ctx,title,username)
	if err!=nil{
		return err 
	}
	return nil 
}

func (d*DBStore)GetMyBooks(ctx context.Context,username string)([]domain.Book,error){
	dBooks,err:=d.BookStorage.GetMyBooks(ctx,username)
	if err!=nil{
		return nil,err
	}
	return dBooks,nil
}
func New(db *sqlx.DB) *DBStore {
	return &DBStore{
		BookStorage: NewBookStorage(db),
	}
}
