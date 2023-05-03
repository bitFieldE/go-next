-- +goose Up
-- +goose StatementBegin
INSERT INTO users (id, name, email, password)
VALUES (1, 'John Doe', 'test@example.com', 'asdfvfvf'),
  (2, 'Joe Do', 'test2@example.com', 'ascasdfvfvf'),
  (3, 'Joe Do', 'test3@example.com', 'assdfvfvf');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE id = 1;
DELETE FROM users WHERE id = 2;
DELETE FROM users WHERE id = 3;
-- +goose StatementEnd
