package entity

import db "github.com/Bakhram74/effective-mobile-test-work.git/db/sqlc"

type Pagination struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalPage int `json:"total_page"`
}

type VerseResponse struct {
	Group      string   `json:"group"`
	Name       string   `json:"name"`
	Verse      []string `json:"verse"`
	Pagination `json:"pagination"`
}


type SongResponse struct {
	Songs      []db.Song
	Pagination `json:"pagination"`
}

type FilteredSongsParams struct {
	Limit     int32
	Offset    int32
	Direction string
	SortValue string
}
