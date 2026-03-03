-- Down migration reverses the Up migration
DROP TABLE IF EXISTS signals;
DROP TYPE IF EXISTS action;
DROP TYPE IF EXISTS symbol;
