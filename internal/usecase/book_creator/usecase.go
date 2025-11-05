package bookcreator

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

func(u*UseCase)CreateBook(ctx context.Context,createBook domain.CreateBook)error{
	err:=u.BookStorage.CreateBook(ctx,createBook)
	if err!=nil{
		return err 
	}
	return nil 
}