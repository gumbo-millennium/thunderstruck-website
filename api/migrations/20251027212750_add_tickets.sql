-- +goose Up
-- +goose StatementBegin
CREATE TABLE tickets (
  id UUID DEFAULT GEN_RANDOM_UUID(),
  type INT NOT NULL,
  email VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,

  PRIMARY KEY (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tickets;
-- +goose StatementEnd
