package service

import (
	"context"

	db "github.com/Bakhram74/effective-mobile-test-work.git/db/sqlc"
)

type Song interface {
	Create(ctx context.Context, params db.CreateSongsParams) (db.Song, error)
}

type Service struct {
	Song
}

func NewService(Queries *db.Queries) Service {
	songService := NewSongService(Queries)
	return Service{
		Song: songService,
	}

}
