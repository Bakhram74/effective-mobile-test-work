-- name: CreateSongs :one
INSERT INTO songs (
"group",
"name",
"releaseDate",
"text",
"link"
) VALUES (
 $1,$2,$3,$4,$5
) RETURNING *;

