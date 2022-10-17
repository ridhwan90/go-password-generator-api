-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserbyUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;

-- name: CreateUser :one
INSERT INTO users (
  id, username, hashed_password, userinfo_uuid
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUsers :one
UPDATE users
set username = $2,
hashed_password = $3
WHERE id = $1
RETURNING *;

-- name: GetUserWithInfo :one
SELECT * from users
LEFT JOIN userinfo ON users.userinfo_uuid = userinfo.id
WHERE users.id = $1;

-- name: UpdateUserInfoUUID :one
UPDATE users
set userinfo_uuid = $2
WHERE id = $1
RETURNING *;