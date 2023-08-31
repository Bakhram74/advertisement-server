-- name: CreateUser :one
INSERT INTO users ("username",
                   "phone_number",
                   "password")
VALUES ($1, $2, $3) RETURNING *;


-- name: PartialUpdateUser :one
UPDATE users
SET username = CASE WHEN @update_username::boolean THEN @username::TEXT ELSE username END,
    phone_number  = CASE WHEN @update_phone_number::boolean THEN @phone_number::TEXT ELSE phone_number END,
    password  = CASE WHEN @update_password::boolean THEN @password::TEXT ELSE password END
WHERE id = @id
RETURNING *;

-- -- name: UpdateUsername :one
-- UPDATE users
-- SET username = $2
-- WHERE id = $1
--     RETURNING *;
--
--
-- -- name: UpdateUserPhone :one
-- UPDATE users
-- SET phone_number = $2
-- WHERE id = $1 RETURNING *;
--
-- -- name: UpdateUser :one
-- UPDATE users
-- SET username     = $2,
--     phone_number = $3
-- WHERE id = $1 RETURNING *;