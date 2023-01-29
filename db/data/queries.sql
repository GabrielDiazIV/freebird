-- name: GetTweets :many
SELECT * FROM tweets WHERE bird_fk = $1;

-- name: GetBirds :many
SELECT * FROM birds ORDER BY score;

-- name: GetBird :one
SELECT * FROM birds WHERE id = $1;
