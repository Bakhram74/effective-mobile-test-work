package service

import (
	"context"

	db "github.com/Bakhram74/effective-mobile-test-work.git/db/sqlc"
	"github.com/Bakhram74/effective-mobile-test-work.git/internal/service/entity"
	"github.com/google/uuid"
)

type Song interface {
	GetAllSongs(ctx context.Context) ([]db.Song, error)
	Create(ctx context.Context, params db.CreateSongParams) (db.Song, error)
	Update(ctx context.Context, params db.UpdateSongParams) (db.Song, error)
	Get(ctx context.Context, id uuid.UUID) (db.Song, error)
	Delete(ctx context.Context, id uuid.UUID) error
	SongVerses(ctx context.Context, params db.SongVersesParams) (int64, []db.SongVersesRow, error)
	FilteredSongs(ctx context.Context, params entity.FilteredSongsParams) (int64, []db.Song, error)
}

type Service struct {
	Song
}

func NewService(store db.Store) Service {
	songService := NewSongService(store)
	return Service{
		Song: songService,
	}

}
