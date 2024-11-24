package service

import (
	"context"

	db "github.com/Bakhram74/effective-mobile-test-work.git/db/sqlc"
	"github.com/Bakhram74/effective-mobile-test-work.git/internal/service/entity"

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

func (s SongService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.queries.DeleteSong(ctx, id)
}

func (s SongService) SongVerses(ctx context.Context, params db.SongVersesParams) (int64, []db.SongVersesRow, error) {
	args := db.CountSongVersesParams{
		Group: params.Group,
		Name:  params.Name,
	}

	count, err := s.queries.CountSongVerses(ctx, args)
	if err != nil {
		return 0, []db.SongVersesRow{}, err
	}

	rows, err := s.queries.SongVerses(ctx, params)
	if err != nil {
		return 0, []db.SongVersesRow{}, err
	}
	return count, rows, nil
}

func (s SongService) FilteredSongs(ctx context.Context, params entity.FilteredSongsParams) (int64, []db.Song, error) {

	if params.SortValue == "date" {
		return 0, nil, nil
	}

	if params.Direction == "desc" {
		args := db.FilteredSongsDescParams{
			Column1: params.SortValue,
			Limit:   params.Limit,
			Offset:  params.Offset,
		}
		rows, err := s.queries.FilteredSongsDesc(ctx, args)
		if err != nil {
			return 0, nil, err
		}
		songs := make([]db.Song, len(rows))
		for i, item := range rows {
			songs[i].ID = item.ID
			songs[i].Group = item.Group
			songs[i].Name = item.Name
			songs[i].ReleaseDate = item.ReleaseDate
			songs[i].Text = item.Text
			songs[i].Link = item.Link
		}
		return rows[0].Total, songs, nil
	}

	args := db.FilteredSongsAscParams{
		Column1: params.SortValue,
		Limit:   params.Limit,
		Offset:  params.Offset,
	}

	rows, err := s.queries.FilteredSongsAsc(ctx, args)
	if err != nil {
		return 0, nil, err
	}

	songs := make([]db.Song, len(rows))
	for i, item := range rows {
		songs[i].ID = item.ID
		songs[i].Group = item.Group
		songs[i].Name = item.Name
		songs[i].ReleaseDate = item.ReleaseDate
		songs[i].Text = item.Text
		songs[i].Link = item.Link
	}

	return rows[0].Total, songs, nil
}
