package service

import (
	"context"

	db "github.com/Bakhram74/effective-mobile-test-work.git/db/sqlc"
	"github.com/google/uuid"
)

type Song interface {
	Create(ctx context.Context, params db.CreateSongParams) (db.Song, error)
	Update(ctx context.Context,params db.UpdateSongParams) (db.Song, error)
	Get(ctx context.Context, id uuid.UUID) (db.Song, error)
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
