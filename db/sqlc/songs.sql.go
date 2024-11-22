// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: songs.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createSong = `-- name: CreateSong :one
INSERT INTO songs (
"group",
"name",
"release_date",
"text",
"link"
) VALUES (
 $1,$2,$3,$4,$5
) RETURNING id, "group", name, release_date, text, link
`

type CreateSongParams struct {
	Group       string `json:"group"`
	Name        string `json:"name"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func (q *Queries) CreateSong(ctx context.Context, arg CreateSongParams) (Song, error) {
	row := q.db.QueryRow(ctx, createSong,
		arg.Group,
		arg.Name,
		arg.ReleaseDate,
		arg.Text,
		arg.Link,
	)
	var i Song
	err := row.Scan(
		&i.ID,
		&i.Group,
		&i.Name,
		&i.ReleaseDate,
		&i.Text,
		&i.Link,
	)
	return i, err
}

const getSong = `-- name: GetSong :one
SELECT id, "group", name, release_date, text, link FROM songs
WHERE id = $1
`

func (q *Queries) GetSong(ctx context.Context, id uuid.UUID) (Song, error) {
	row := q.db.QueryRow(ctx, getSong, id)
	var i Song
	err := row.Scan(
		&i.ID,
		&i.Group,
		&i.Name,
		&i.ReleaseDate,
		&i.Text,
		&i.Link,
	)
	return i, err
}

const updateSong = `-- name: UpdateSong :one
UPDATE songs
SET
  "group" = COALESCE($1, "group"),
  name = COALESCE($2, name),
  release_date = COALESCE($3, release_date),
  text = COALESCE($4, text),
  link = COALESCE($5, link)
WHERE id = $6
RETURNING id, "group", name, release_date, text, link
`

type UpdateSongParams struct {
	Group       string    `json:"group"`
	Name        string    `json:"name"`
	ReleaseDate string    `json:"release_date"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
	ID          uuid.UUID `json:"id"`
}

func (q *Queries) UpdateSong(ctx context.Context, arg UpdateSongParams) (Song, error) {
	row := q.db.QueryRow(ctx, updateSong,
		arg.Group,
		arg.Name,
		arg.ReleaseDate,
		arg.Text,
		arg.Link,
		arg.ID,
	)
	var i Song
	err := row.Scan(
		&i.ID,
		&i.Group,
		&i.Name,
		&i.ReleaseDate,
		&i.Text,
		&i.Link,
	)
	return i, err
}
