-- name: CreateUser :one
INSERT INTO users ("username",
                   "phone_number",
                   "password")
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateUsername :one
UPDATE users
SET username = $2
WHERE id = $1
    RETURNING *;


-- name: UpdateUserPhone :one
UPDATE users
SET phone_number = $2
WHERE id = $1 RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET username     = $2,
    phone_number = $3
WHERE id = $1;