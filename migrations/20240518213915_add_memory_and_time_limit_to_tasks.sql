-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
ALTER TABLE tasks ADD COLUMN time_limit INTEGER NOT NULL default 1000;
ALTER TABLE tasks ADD COLUMN memory_limit INTEGER NOT NULL default 1024;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
ALTER TABLE tasks DROP COLUMN time_limit;
ALTER TABLE tasks DROP COLUMN memory_limit;
