CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "username" text NOT NULL,
                         "phone_number" text NOT NULL,
                         "password" text NOT NULL,
                         "role" text DEFAULT (user),
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);
