// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CountSongVerses(ctx context.Context, arg CountSongVersesParams) (int64, error)
	CreateSong(ctx context.Context, arg CreateSongParams) (Song, error)
	DeleteSong(ctx context.Context, id uuid.UUID) error
	FilteredSongsAsc(ctx context.Context, arg FilteredSongsAscParams) ([]FilteredSongsAscRow, error)
	FilteredSongsDesc(ctx context.Context, arg FilteredSongsDescParams) ([]FilteredSongsDescRow, error)
	GetAllSong(ctx context.Context) ([]Song, error)
	GetSong(ctx context.Context, id uuid.UUID) (Song, error)
	SongVerses(ctx context.Context, arg SongVersesParams) ([]SongVersesRow, error)
	UpdateSong(ctx context.Context, arg UpdateSongParams) (Song, error)
}

var _ Querier = (*Queries)(nil)
