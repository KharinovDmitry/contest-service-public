-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS task_tag_pairs (
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    task_id BIGINT NOT NULL,
    tag_id BIGINT NOT NULL,

    FOREIGN KEY (task_id) REFERENCES tasks(id)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE task_tag_pairs;
