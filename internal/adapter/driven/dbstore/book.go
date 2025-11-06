package dbstore

import (
	"context"
	"errors"
	"fmt"


	"time"

	"github.com/Ilja-R/library-service/internal/domain"
	"github.com/Ilja-R/library-service/internal/errs"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"

	"os"
)

type BookStorage struct {
	db *sqlx.DB
}

func NewBookStorage(db *sqlx.DB) *BookStorage {
	return &BookStorage{db: db}
}

type Book struct {
	ID int `db:"id"`
	Title string `db:"title"`
	Pub_date time.Time `db:"pub_date"`
	Publisher string `db:"publisher"`
	Genre string `db:"genre"`
	Pages int `db:"pages"`
	Description string `db:"description"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	
}
type CreateBook struct{
	ID int `db:"id"`
	Title string `db:"title"`
	AuthorName string `db:"name"`
	AuthorSurname string `db:"surname"`
	Pub_date time.Time `db:"pub_date"`
	Publisher string `db:"publisher"`
	Genre string `db:"genre"`
	Pages int `db:"pages"`
	Description string `db:"description"`

}

type Author struct{
	ID int `db:"id"`
	Name string `db:"name"`
	Surname string `db:"surname"`
}

func (b *Book) FromDomain(dBook domain.Book) {
	b.ID=dBook.ID
	b.Title=dBook.Title
	b.Pub_date=dBook.Pub_date
	b.Publisher=dBook.Publisher
	b.Genre=dBook.Genre
	b.Pages=dBook.Pages
	b.Description=dBook.Description
}

func (b *Book) ToDomain() *domain.Book {
	return &domain.Book{
		ID:b.ID,
		Title:b.Title,
		Pub_date:b.Pub_date,
		Publisher: b.Publisher,
		Genre:b.Genre,
		Pages:b.Pages,
		Description:b.Description,
		Created_at: b.CreatedAt,
		Updated_at: b.UpdatedAt,
	}
}

func (b *BookStorage) GetAllBooks(ctx context.Context) (books []domain.Book, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "dbstore.GetAllBooks").Logger()

	var dbBooks []Book
	if err = b.db.SelectContext(ctx, &dbBooks, `
		SELECT id,title,pub_date,publisher,genre,pages,description,created_at,updated_at
		FROM books
		`); err != nil {
		logger.Err(err).Msg("error selecting products")
		return nil, b.translateError(err)
	}
	
	
	for _, dbBook := range dbBooks{
		books = append(books,*dbBook.ToDomain())
	}
	
	return books, nil
}

func (b *BookStorage) GetBookByID(ctx context.Context, id int) (domain.Book, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "dbstore.GetBookByID").Logger()

	var dbBook Book
	if err := b.db.GetContext(ctx, &dbBook, `
		SELECT id,title,pub_date,publisher,genre,pages,description,created_at,updated_at
		FROM books
		WHERE id = $1`, id); err != nil {
		logger.Err(err).Msg("select error")
		return domain.Book{}, b.translateError(err)
	}

	return *dbBook.ToDomain(), nil
}
type CompareBooks struct {
	Name string `db:"name"`
	Surname string `db:"surname"`
	Title string `db:"title"`

}
func (b *BookStorage) CreateProduct(ctx context.Context, createBook domain.CreateBook) (error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "dbstore.CreateBook").Logger()
	logger.Debug().Any("createBook",createBook).Send()
	author:=Author{}
	err:=b.db.GetContext(ctx,&author,`
	SELECT id,name,surname from authors
	where name =$1 and surname =$2
	`,createBook.AuthorName,createBook.AuthorSurname)
	if err!=nil{
		logger.Err(err).Send()
		if errors.Is(err,errs.ErrNoFollowingAuthor){
			return errs.ErrNoFollowingAuthor
		}
		return err
	}
	logger.Info().Any("author",author).Send()
	compare:=CompareBooks{}
	err=b.db.GetContext(ctx,&compare,`
	select name,surname,title from book_author 
	join books on book_id = books.id
    join authors on author_id = authors.id
	where name=$1 and surname =$2 and title = $3
	`,author.Name,author.Surname,createBook.Title)
	logger.Debug().Any("ok",compare).Send()
	if err!=nil{
		logger.Err(err).Send()
	}
	if compare.Name==createBook.AuthorName&&compare.Surname==createBook.AuthorSurname&&compare.Title==createBook.Title{
		return fmt.Errorf("%s","following combination of name surname and title already exists")
	}
	


	_,err=b.db.ExecContext(ctx,`
	INSERT INTO books (title,pub_date,publisher,genre,pages,description) VALUES ($1,$2,$3,$4,$5,$6)
	`,createBook.Title,createBook.Pub_date,createBook.Publisher,createBook.Genre,createBook.Pages,createBook.Description)
	if err!=nil{
		fmt.Println(err)
		return err
	}
	var id int 
	err=b.db.GetContext(ctx,&id,`
	SELECT id FROM books ORDER BY id DESC LIMIT 1 ;
	`)
	if err!=nil{
		fmt.Println(err)
		return err
	}
	_,err=b.db.ExecContext(ctx,`
	INSERT INTO book_author (book_id,author_id) VALUES ($1,$2)
	`,id,author.ID)
	if err != nil {
		fmt.Println(err)
		return b.translateError(err)
	}

	return nil
}

type Book_Author struct{
	ID int `db:"id"`
	Name string `db:"name"`
	Surname string `db:"surname"`

}

func (b *BookStorage) UpdateBookByID(ctx context.Context, UpdateBook domain.UpdateBookBody,id int) (err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "dbstore.UpdateBookByID").Logger()
	var bookAuthor Book_Author
 	err= b.db.GetContext(ctx, &bookAuthor,`
	select book_author.id,name,surname from book_author 
	inner join books on book_author.book_id=books.id 
	inner join authors on book_author.author_id =authors.id 
	where books.id=$1
	`, id)
	if err!=nil{
		logger.Err(err).Send()
		return err

	}
	logger.Debug().Any("bookAuhtor",bookAuthor).Send()
	logger.Debug().Any("Updatebook",UpdateBook).Send()
	if bookAuthor.Name!=UpdateBook.AuthorName||bookAuthor.Surname!=UpdateBook.AuthorSurname{
		var authorID int 
		err=b.db.GetContext(ctx,&authorID,`
		select id from authors where name =$1 and surname =$2
		`,UpdateBook.AuthorName,UpdateBook.AuthorSurname)
		if authorID ==0{
			return errs.ErrNoFollowingAuthor
		}
		if err!=nil{
			logger.Err(err).Send()
			return err

		}
		_,err=b.db.ExecContext(ctx,`
		update books set title=$1,publisher=$2,genre=$3,pages=$4,description=$5
		where books.id=$6
		`,UpdateBook.Title,UpdateBook.Publisher,
		  UpdateBook.Genre,UpdateBook.Pages,
		  UpdateBook.Description,id)
		
		if err!=nil{
			logger.Err(err).Send()
			return err

		}
		_,err=b.db.ExecContext(ctx,`
		update book_author set author_id=$1 where 
		book_id=$2`,authorID,id)
		if err!=nil{
			logger.Err(err).Send()
			return err
		}
	}else{
		_,err=b.db.ExecContext(ctx,`
		update books set title=$1,publisher=$2,genre=$3,pages=$4,description=$5
		where books.id=$6
		`,UpdateBook.Title,UpdateBook.Publisher,
		  UpdateBook.Genre,UpdateBook.Pages,
		  UpdateBook.Description,id)
		
		if err!=nil{
			logger.Err(err).Send()
			return err

		}
	}

	
	
	return nil
}

func (b *BookStorage) DeleteBookByID(ctx context.Context, id int) (error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "dbstore.DeleteBookByID").Logger()
	dbBook,_:=b.GetBookByID(ctx,id )
	if dbBook.ID==0{
		return errs.ErrBookNotFound
	}
	_, err := b.db.ExecContext(ctx, `DELETE FROM book_author WHERE book_id = $1`, id)
	if err != nil {
		logger.Err(err).Msg("error deleting book from book_author")
		return b.translateError(err)
	}
	_, err = b.db.ExecContext(ctx, `DELETE FROM book_username WHERE book_id = $1`, id)
	if err != nil {
		logger.Err(err).Msg("error deleting book from book_author")
		return b.translateError(err)
	}
	_, err = b.db.ExecContext(ctx, `DELETE FROM book_author_username  WHERE id = $1`, id)
	if err != nil {
		logger.Err(err).Msg("error deleting book from books")
		return b.translateError(err)
	}
	_, err = b.db.ExecContext(ctx, `DELETE FROM books WHERE id = $1`, id)
	if err != nil {
		logger.Err(err).Msg("error deleting book from books")
		return b.translateError(err)
	}
	

	return nil
}

func(b *BookStorage)SearchByTitle(ctx context.Context, title string)(books []domain.Book,err error){
	dbBooks :=make([]Book,0)
	err=b.db.SelectContext(ctx,&books,`
	SELECT id,title,pub_date,publisher,genre,pages,description,created_at,updated_at FROM books
	where title =$1
	`,title)
	if err!=nil{
		return nil ,err 
	}
	for _, dbBook := range dbBooks{
		books = append(books, *dbBook.ToDomain())
	}
	return books,nil
}