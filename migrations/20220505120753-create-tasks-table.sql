
-- +migrate Up
CREATE TABLE IF NOT EXISTS tasks
(
    id          VARCHAR PRIMARY KEY,
    name        VARCHAR NOT NULL,
    description TEXT,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL
);
-- +migrate Down
DROP TABLE tasks;