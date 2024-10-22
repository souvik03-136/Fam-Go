-- youtube_queries.sql

-- name: CreateVideo :exec
INSERT INTO videos (title, description, published_at, thumbnail_url, video_url)
VALUES (?, ?, ?, ?, ?);

-- name: GetVideoByID :one
SELECT id, title, description, published_at, thumbnail_url, video_url
FROM videos
WHERE id = ?;

-- name: ListVideos :many
SELECT id, title, description, published_at, thumbnail_url, video_url
FROM videos
ORDER BY published_at DESC
LIMIT ? OFFSET ?;

-- name: UpdateVideo :exec
UPDATE videos
SET title = ?, description = ?, published_at = ?, thumbnail_url = ?, video_url = ?
WHERE id = ?;

-- name: DeleteVideo :exec
DELETE FROM videos
WHERE id = ?;

-- name: GetLastInsertedVideo :one
SELECT id, title, description, published_at, thumbnail_url, video_url 
FROM videos 
WHERE id = (SELECT LAST_INSERT_ID());  -- Fetch the newly inserted video details
