-- name: CreateSong :one
INSERT INTO songs (
"group",
"name",
"release_date",
"text",
"link"
) VALUES (
 $1,$2,$3,$4,$5
) RETURNING *;

-- name: GetSong :one
SELECT * FROM songs
WHERE id = $1;

-- name: DeleteSong :exec
DELETE FROM songs
WHERE id = $1;



-- name: UpdateSong :one
UPDATE songs
SET
  "group" = COALESCE($1, "group"),
  name = COALESCE($2, name),
  release_date = COALESCE($3, release_date),
  text = COALESCE($4, text),
  link = COALESCE($5, link)
WHERE id = $6
RETURNING *;
