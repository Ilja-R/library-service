package bookdeleter

import (
	"context"

	"github.com/Ilja-R/library-service/internal/config"
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

func(u*UseCase)DeleteBookByID(ctx context.Context,id int)error{
	err:=u.BookStorage.DeleteBookByID(ctx,id)
	if err!=nil{
		return err 
	}
	return nil 
}