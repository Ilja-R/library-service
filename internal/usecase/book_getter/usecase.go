package bookgetter

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Ilja-R/library-service/internal/config"
	"github.com/Ilja-R/library-service/internal/domain"
	"github.com/Ilja-R/library-service/internal/errs"
	"github.com/Ilja-R/library-service/internal/port/driven"
	"github.com/redis/go-redis/v9"
)

const DEFAULT_BOOK_TTL=5*time.Minute

type UseCase struct{
	cfg *config.Config
	BookStorage driven.BookStorage
	Cache driven.Cache
}


func New (cfg *config.Config,bookStorage driven.BookStorage,cache driven.Cache )*UseCase{
	return &UseCase{
		cfg:cfg,
		BookStorage: bookStorage,
		Cache:cache,
	}
}

func (u*UseCase)GetAllBooks(ctx context.Context)([]domain.Book,error){
	books,err:=u.BookStorage.GetAllBooks(ctx)
	if err!=nil{
		return nil,err
	}
	return books,nil
}


func(u*UseCase)GetBookByID(ctx context.Context,id int )(book domain.Book,err error){
	err=u.Cache.Get(ctx,fmt.Sprintf("book_%d",id),&book)
	if err==nil{
		return book ,nil
	}
	if !errors.Is(err,redis.Nil){
		return domain.Book{},err
	}
	book,err=u.BookStorage.GetBookByID(ctx,id)
	if err!=nil{
		if errors.Is(err,errs.ErrNotfound){
			return domain.Book{},errs.ErrBookNotFound
		}else{
			return domain.Book{},err
		}
	}
	
	err=u.Cache.Set(ctx,fmt.Sprintf("book_%d",id),book,DEFAULT_BOOK_TTL)
	if err!=nil{
		return domain.Book{},err
	}
	return book,nil 

}