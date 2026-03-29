package usecase

import (
	"context"

)

type BookDeleter interface {
	DeleteBookByID(ctx context.Context,id int)error
}