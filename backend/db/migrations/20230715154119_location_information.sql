-- +goose Up
-- +goose StatementBegin
CREATE TABLE location_information (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id INTEGER REFERENCES users(id) NOT NULL,
  location_id INTEGER REFERENCES locations(id) NOT NULL,
  content CHARACTER VARYING(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
CREATE INDEX idx_user_id ON location_information (user_id);
CREATE INDEX idx_location_id ON location_information (location_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE location_information;
-- +goose StatementEnd
