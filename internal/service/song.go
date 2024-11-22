package service

import (
	"context"

	db "github.com/Bakhram74/effective-mobile-test-work.git/db/sqlc"

	"github.com/google/uuid"
)

type SongService struct {
	queries *db.Queries
}

func NewSongService(queries *db.Queries) *SongService {
	return &SongService{
		queries: queries,
	}
}

func (s SongService) Create(ctx context.Context, params db.CreateSongParams) (db.Song, error) {
	return s.queries.CreateSong(ctx, params)
}

func (s SongService) Get(ctx context.Context, id uuid.UUID) (db.Song, error) {
	return s.queries.GetSong(ctx, id)
}

func (s SongService) Update(ctx context.Context, params db.UpdateSongParams) (db.Song, error) {
	return s.queries.UpdateSong(ctx, params)
}
