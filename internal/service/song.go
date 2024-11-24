package service

import (
	"context"
	"fmt"
	"sort"
	"strings"

	db "github.com/Bakhram74/effective-mobile-test-work.git/db/sqlc"
	"github.com/Bakhram74/effective-mobile-test-work.git/internal/service/entity"

	"github.com/google/uuid"
)

type SongService struct {
	store db.Store
}

func NewSongService(store db.Store) *SongService {
	return &SongService{
		store: store,
	}
}

func (s SongService) GetAllSongs(ctx context.Context) ([]db.Song, error) {
	return s.store.GetAllSong(ctx)
}

func (s SongService) Create(ctx context.Context, params db.CreateSongParams) (db.Song, error) {
	return s.store.CreateSong(ctx, params)
}

func (s SongService) Get(ctx context.Context, id uuid.UUID) (db.Song, error) {
	return s.store.GetSong(ctx, id)
}

func (s SongService) Update(ctx context.Context, params db.UpdateSongParams) (db.Song, error) {
	return s.store.UpdateSong(ctx, params)
}

func (s SongService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.store.DeleteSong(ctx, id)
}

func (s SongService) SongVerses(ctx context.Context, params db.SongVersesParams) (int64, []db.SongVersesRow, error) {
	return s.store.SongVersesTX(ctx, params)
}

func (s SongService) FilteredSongs(ctx context.Context, params entity.FilteredSongsParams) (int64, []db.Song, error) {

	if params.SortValue == "date" {
		count, songs, err := s.sortByDate(ctx, params)
		if err != nil {
			return 0, nil, err
		}
		return count, songs, nil
	}

	if params.Direction == "desc" {
		args := db.FilteredSongsDescParams{
			Column1: params.SortValue,
			Limit:   params.Limit,
			Offset:  params.Offset,
		}
		rows, err := s.store.FilteredSongsDesc(ctx, args)
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

	rows, err := s.store.FilteredSongsAsc(ctx, args)
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

func (s SongService) sortByDate(ctx context.Context, params entity.FilteredSongsParams) (int64, []db.Song, error) {

	songs, err := s.store.GetAllSong(ctx)
	if err != nil {
		return 0, nil, err
	}

	newSongs := make([]db.Song, len(songs))

	for i, song := range songs {
		split := strings.Split(song.ReleaseDate, ".")
		reversed := strings.Join([]string{split[2], split[1], split[0]}, "")
		newSongs[i].ID = song.ID
		newSongs[i].Group = song.Group
		newSongs[i].Name = song.Name
		newSongs[i].ReleaseDate = reversed
		newSongs[i].Text = song.Text
		newSongs[i].Link = song.Link
	}
	if params.Direction == "desc" {
		sort.SliceStable(newSongs, func(i, j int) bool {
			return newSongs[i].ReleaseDate > newSongs[j].ReleaseDate
		})
	} else {
		sort.SliceStable(newSongs, func(i, j int) bool {
			return newSongs[i].ReleaseDate < newSongs[j].ReleaseDate
		})
	}

	for i, song := range newSongs {
		song.ReleaseDate = fmt.Sprintf("%s.%s.%s", song.ReleaseDate[6:], song.ReleaseDate[4:6], song.ReleaseDate[:4])
		newSongs[i] = song
	}

	limit := int(params.Limit + params.Offset)
	if limit > len(songs) {
		limit = len(songs)
	}

	if int(params.Offset) > len(songs) {
		return int64(len(songs)), nil, nil
	}

	return int64(len(songs)), newSongs[params.Offset:limit], nil
}
