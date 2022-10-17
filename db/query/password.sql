-- name: GetPassword :one
SELECT * FROM password
WHERE id = $1 LIMIT 1;

-- name: ListPassword :many
SELECT * FROM password
ORDER BY user_uuid;

-- name: ListPasswordbyUser :many
SELECT * FROM password
WHERE user_uuid = $1;

-- name: CreatePassword :one
INSERT INTO password (
  id, user_uuid, site, site_username, site_email, generated_password
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: DeletePassword :exec
DELETE FROM password
WHERE id = $1;

-- name: UpdatePassword :one
UPDATE password
set site = $2,
site_username = $3,
site_email = $4,
generated_password = $5
WHERE id = $1
RETURNING *;

-- name: GetUserPassword :many
SELECT * FROM password
WHERE user_uuid = $1;