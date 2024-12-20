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

-- name: SongVerses :many
SELECT 
  verses."name",
  verses.verse
FROM (
  SELECT 
    songs."name",
    unnest(string_to_array(songs."text", E'\n')) AS verse
  FROM songs
  WHERE songs."group" = $1 AND songs."name" = $2
) AS verses
WHERE verses.verse <> ''
LIMIT $3 OFFSET $4;

-- name: CountSongVerses :one
SELECT 
  COUNT(*) AS verse_count
FROM (
  SELECT 
    unnest(string_to_array("text", E'\n')) AS verse
  FROM songs
  WHERE "group" = $1 AND "name" = $2
) AS verses
WHERE verse <> '';

-- name: FilteredSongsAsc :many
WITH total_count AS (
  SELECT COUNT(*) AS total
  FROM songs
)
SELECT 
  songs.*, 
  total_count.total
FROM songs, total_count
ORDER BY 
  CASE 
    WHEN $1 = 'name' THEN "name"
    WHEN $1 = 'group' THEN "group"
    WHEN $1 = 'text' THEN "text"
    WHEN $1 = 'link' THEN "link"
  END ASC
LIMIT $2 OFFSET $3;

-- name: FilteredSongsDesc :many
WITH total_count AS (
  SELECT COUNT(*) AS total
  FROM songs
)
SELECT 
  songs.*, 
  total_count.total
FROM songs, total_count
ORDER BY 
  CASE 
    WHEN $1 = 'name' THEN "name"
    WHEN $1 = 'group' THEN "group"
    WHEN $1 = 'text' THEN "text"
    WHEN $1 = 'link' THEN "link"
  END DESC
LIMIT $2 OFFSET $3;

-- name: GetAllSong :many
SELECT * FROM songs;
