-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS tasks(
    id BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL UNIQUE,
    title VARCHAR(256) NOT NULL,
    text TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS launches (
    id BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL UNIQUE,
    user_id BIGINT NOT NULL,
    task_id BIGINT NOT NULL,
    date DATE NOT NULL,
    code TEXT NOT NULL,
    test_result_code VARCHAR(2) NOT NULL,
    description TEXT NOT NULL,
    points INT NOT NULL,

    FOREIGN KEY (task_id) REFERENCES tasks(id)
);


CREATE TABLE IF NOT EXISTS tests(
    id BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL UNIQUE,
    task_id BIGINT NOT NULL,
    input VARCHAR(512) NOT NULL,
    expected_result VARCHAR(512) NOT NULL,
    points INTEGER NOT NULL,

    FOREIGN KEY (task_id) REFERENCES tasks(id)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE tests;
DROP TABLE tasks;
DROP TABLE launches;
