// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"github.com/google/uuid"
)

type Song struct {
	ID          uuid.UUID `json:"id"`
	Group       string    `json:"group"`
	Name        string    `json:"name"`
	ReleaseDate string    `json:"releaseDate"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
}
