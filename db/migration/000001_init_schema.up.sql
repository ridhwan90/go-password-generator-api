CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL UNIQUE,
  "hashed_password" varchar NOT NULL,
  "userinfo_uuid" uuid,
  "password_change_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "userinfo" (
  "id" uuid PRIMARY KEY,
  "first_name" varchar,
  "last_name" varchar,
  "phone_number" varchar,
  "email" varchar UNIQUE,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "password" (
  "id" uuid PRIMARY KEY,
  "user_uuid" uuid,
  "site" varchar NOT NULL UNIQUE,
  "site_username" varchar,
  "site_email" varchar,
  "generated_password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "users" ADD FOREIGN KEY ("userinfo_uuid") REFERENCES "userinfo" ("id") ON DELETE CASCADE;

ALTER TABLE "password" ADD FOREIGN KEY ("user_uuid") REFERENCES "users" ("id") ON DELETE CASCADE;