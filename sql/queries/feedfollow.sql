-- name: CreateFeedFollow :one
WITH inserted_feed_follows AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT
    users.name AS username,
    feeds.name AS feed_name,
    inserted_feed_follows.*
FROM
    inserted_feed_follows
    INNER JOIN users ON users.id = inserted_feed_follows.user_id
    INNER JOIN feeds ON feeds.id = inserted_feed_follows.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT
    users.name AS username,
    feeds.name AS feed_name,
    feed_follows.*
FROM
    users
    INNER JOIN feed_follows ON users.id = feed_follows.user_id
    INNER JOIN feeds ON feeds.id = feed_follows.feed_id
WHERE
    users.name = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE
    feed_follows.user_id = $1
    AND feed_follows.feed_id = $2;
