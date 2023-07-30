-- +goose Up
-- +goose StatementBegin
CREATE TABLE locations (
  id SERIAL PRIMARY KEY NOT NULL,
  location_place_id INTEGER REFERENCES location_places(id) NOT NULL,
  name CHARACTER VARYING(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
CREATE INDEX idx_location_place_id ON locations (location_place_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE locations;
-- +goose StatementEnd
