-- +goose Up
-- +goose StatementBegin
CREATE TABLE students(
    id SERIAL PRIMARY KEY,
    code VARCHAR(4) NOT NULL,
    name VARCHAR(150) NOT NULL
);

CREATE TABLE scores(
    id SERIAL PRIMARY KEY,
    exercise_date TIMESTAMP NOT NULL,
    score INT NOT NULL,
    student_id INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS scores;
DROP TABLE IF EXISTS students;
-- +goose StatementEnd
