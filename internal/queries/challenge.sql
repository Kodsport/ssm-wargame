
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

-- name: ListMonthlyChallenges :many
SELECT * FROM monthly_challenges;

-- name: InsertMonthlyChallenge :exec
INSERT INTO monthly_challenges (challenge_id, start_date, end_date, display_month) VALUES ($1, $2, $3, $4);

-- name: DeleteMonthlyChallenge :exec
DELETE FROM monthly_challenges WHERE challenge_id = $1;

-- name: ChallengeExists :one
SELECT EXISTS(SELECT 1 FROM challenges WHERE id = $1);

-- name: InsertFile :exec
INSERT INTO challenge_files (id, challenge_id, friendly_name, bucket, key, md5, uploaded) VALUES (@id::uuid, @challenge_id::uuid, @fname::text, @bucket::text, @key::text, @md5::text, false);

-- name: FileMarkUploaded :exec
UPDATE challenge_files SET uploaded = true, updated_at = NOW() WHERE id = $1;

-- name: DeleteFile :exec
DELETE FROM challenge_files WHERE id = $1;