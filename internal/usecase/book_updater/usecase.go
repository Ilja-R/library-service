package bookupdater

import (
	"context"

	"github.com/Ilja-R/library-service/internal/config"
	"github.com/Ilja-R/library-service/internal/domain"
	"github.com/Ilja-R/library-service/internal/port/driven"
)
type UseCase struct{
	cfg *config.Config
	BookStorage driven.BookStorage
}

func New(cfg *config.Config, b driven.BookStorage)*UseCase{
	return &UseCase{
		cfg:cfg,
		BookStorage: b,
	}
}

func(u*UseCase)UpdateBook(ctx context.Context,book domain.UpdateBookBody,id int)error{
	err:=u.BookStorage.UpdateBookByID(ctx,book,id )
	if err!=nil{
		return err 
	}
	return nil 
}