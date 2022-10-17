-- name: CreateUserInfo :one
INSERT INTO userinfo (
    id,
    first_name,
    last_name,
    phone_number,
    email
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;