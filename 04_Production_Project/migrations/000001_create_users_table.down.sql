-- Down migration reverses the Up migration
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS user_status;
