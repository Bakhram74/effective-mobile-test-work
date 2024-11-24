package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	SongVersesTX(ctx context.Context, params SongVersesParams) (int64, []SongVersesRow, error)
	Querier
}

type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
