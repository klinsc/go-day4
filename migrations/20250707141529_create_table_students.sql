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
    value BIGINT NOT NULL,
    student_id BIGINT NOT NULL REFERENCES students(id) ON DELETE CASCADE
);

CREATE INDEX idx_scores_student_id ON scores(student_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS scores;
DROP TABLE IF EXISTS students;
-- +goose StatementEnd
