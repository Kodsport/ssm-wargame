-- name: UserIDByDiscordID :one
SELECT id FROM users WHERE discord_id = @discord_id::text;

-- name: InsertUserDiscord :exec
INSERT INTO users (id, discord_id, email) VALUES (@id::uuid, @discord_id::text, @email::text);