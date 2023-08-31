CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "username" text NOT NULL,
                         "phone_number" text NOT NULL,
                         "password" text NOT NULL,
                         "role" text NOT NULL DEFAULT ('user'),
                         "is_banned" bool NOT NULL DEFAULT (false),
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);
