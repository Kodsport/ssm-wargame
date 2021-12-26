
-- name: ListChallengesWithSolves :many
SELECT c.*, COUNT(us.user_id) num_solves 
FROM challenges c LEFT JOIN user_solves us ON us.challenge_id = c.id 
WHERE (NOT c.published = @show_unpublished::bool OR c.published = true)
GROUP BY c.id;

-- name: FlagExists :one
SELECT EXISTS(SELECT 1 FROM flags WHERE challenge_id = $1 AND flag = $2);

-- name: InsertAttempt :exec
INSERT INTO submissions (user_id, challenge_id, successful, input) VALUES ($1, $2, $3, $4);

-- name: InsertSolve :exec
INSERT INTO user_solves (user_id, challenge_id) VALUES ($1, $2);

-- name: UserHasSolved :one
SELECT EXISTS (SELECT * FROM user_solves WHERE challenge_id = $1 AND user_id = $2);

-- name: InsertChallenge :exec
INSERT INTO challenges (id, title, slug, description, score, published) VALUES ($1, $2, $3, $4, $5, $6);