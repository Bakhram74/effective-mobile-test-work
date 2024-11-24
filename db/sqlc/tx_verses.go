package db

import (
	"context"
)

func (store *SQLStore) SongVersesTX(ctx context.Context, params SongVersesParams) (int64, []SongVersesRow, error) {

	var versesRow []SongVersesRow
	var count int64

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		versesRow, err = q.SongVerses(ctx, SongVersesParams{
			Group:  params.Group,
			Name:   params.Name,
			Limit:  params.Limit,
			Offset: params.Offset,
		})
		if err != nil {
			return err
		}

		count, err = q.CountSongVerses(ctx, CountSongVersesParams{
			Group: params.Group,
			Name:  params.Name,
		})

		return err

	})
	return count, versesRow, err
}
