-- +goose Up
-- +goose StatementBegin
CREATE TYPE TICKET_TYPE AS ENUM (
  'entry',
  'crew'
);

CREATE TYPE TICKET_STATE AS ENUM (
  'unused',
  'used'
);

CREATE TABLE tickets (
  id UUID DEFAULT GEN_RANDOM_UUID(),
  type TICKET_TYPE NOT NULL,
  state TICKET_STATE NOT NULL DEFAULT 'unused',
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

DROP TYPE TICKET_TYPE;
DROP TYPE TICKET_STATE;
-- +goose StatementEnd
