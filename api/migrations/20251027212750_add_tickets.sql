-- +goose Up
-- +goose StatementBegin
CREATE TYPE ticket_type AS ENUM (
  'Entry',
  'Crew'
);

CREATE TABLE tickets (
  id UUID DEFAULT GEN_RANDOM_UUID(),
  type TICKET_TYPE NOT NULL,
  value VARCHAR(255) NOT NULL,
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
DROP TYPE ticket_type;
-- +goose StatementEnd
