package dbstore

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/Ilja-R/library-service/internal/domain"
	
	"os"
	
)

type BookStorage struct {
	db *sqlx.DB
}

func NewBookStorage(db *sqlx.DB) *BookStorage {
	return &BookStorage{db: db}
}

type Book struct {
	
}

func (b *Book) FromDomain(dbBook domain.Book) {
	
}

func (b *Book) ToDomain() domain.Book {
	return domain.Book{
		
	}
}

func (b *BookStorage) GetAllBooks(ctx context.Context) (books []domain.Book, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "dbstore.GetAllBooks").Logger()

	var dbBooks []Book
	if err = b.db.SelectContext(ctx, &dbBooks, `
		SELECT id, product_name, manufacturer, product_count, price, created_at, updated_at
		FROM products
		ORDER BY id`); err != nil {
		logger.Err(err).Msg("error selecting products")
		return nil, b.translateError(err)
	}

	for _, dbBook := range dbBooks{
		books = append(books, dbBook.ToDomain())
	}

	return books, nil
}

func (b *BookStorage) GetBookByID(ctx context.Context, id int) (domain.Book, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "dbstore.GetBookByID").Logger()

	var dbBook Book
	if err := b.db.GetContext(ctx, &dbBook, `
		SELECT id, product_name, manufacturer, product_count, price, created_at, updated_at
		FROM products
		WHERE id = $1`, id); err != nil {
		logger.Err(err).Msg("error selecting product")
		return domain.Book{}, b.translateError(err)
	}

	return dbBook.ToDomain(), nil
}

func (b *BookStorage) CreateProduct(ctx context.Context, book domain.Book) (err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "dbstore.CreateBook").Logger()

	var dbBook Book
	dbBook.FromDomain(book)
	_, err = b.db.ExecContext(ctx, `INSERT INTO books ()
					VALUES ($1, $2, $3, $4)`,
		
		
		
		
				)
	if err != nil {
		logger.Err(err).Msg("error inserting book")
		return b.translateError(err)
	}

	return nil
}

func (b *BookStorage) UpdateProductByID(ctx context.Context, book domain.Book) (err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "dbstore.UpdateBookByID").Logger()
	var dbBook Book
	dbBook.FromDomain(book)
	_, err = b.db.ExecContext(ctx, `
		UPDATE books SET WHERE id = `,
		)
	if err != nil {
		logger.Err(err).Msg("error updating book")
		return b.translateError(err)
	}

	return nil
}

func (b *BookStorage) DeleteBookByID(ctx context.Context, id int) (err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "dbstore.DeleteBookByID").Logger()
	_, err = b.db.ExecContext(ctx, `DELETE FROM books WHERE id = $1`, id)
	if err != nil {
		logger.Err(err).Msg("error deleting book")
		return b.translateError(err)
	}

	return nil
}