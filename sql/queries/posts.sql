-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostsForUser :many
WITH followed AS (
SELECT
    users.name AS username,
    feeds.name AS feed_name,
    feed_follows.*
FROM
    users
    INNER JOIN feed_follows ON users.id = feed_follows.user_id
    INNER JOIN feeds ON feeds.id = feed_follows.feed_id
WHERE
    users.name = $1
)

SELECT
    username,
    feed_name,
    posts.*
FROM
    posts
    INNER JOIN followed ON posts.feed_id = followed.feed_id
ORDER BY
    posts.published_at DESC NULLS LAST
LIMIT
    $2;
