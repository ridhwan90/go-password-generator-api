ALTER TABLE IF EXISTS "users" DROP CONSTRAINT IF EXISTS "users_userinfo_uuid_fkey";
ALTER TABLE IF EXISTS "password" DROP CONSTRAINT IF EXISTS "password_user_uuid_fkey";

DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "userinfo";
DROP TABLE IF EXISTS "password";