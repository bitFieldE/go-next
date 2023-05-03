-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id serial PRIMARY KEY,
  name  character varying(100) NOT NULL,
  email character varying(255) NOT NULL,
  password character varying(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd

