package usecase

import "context"

type BookOrderer interface {
	OrderBookByTitle(ctx context.Context,title string, username string)error
}