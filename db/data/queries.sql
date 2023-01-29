-- name: GetTweets :many
SELECT * FROM tweets WHERE bird_fk = $1;

-- name: GetBirds :many
SELECT * FROM birds ORDER BY score;
