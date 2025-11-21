-- +goose Up
-- +goose StatementBegin
CREATE TYPE ORDER_STATE AS ENUM (
  'pending',
  'paid',
  'cancelled'
);

CREATE TABLE orders (
  id UUID DEFAULT GEN_RANDOM_UUID(),
  ticket_id UUID,
  reference VARCHAR(255) NOT NULL,
  checkout VARCHAR(255) NOT NULL,
  state ORDER_STATE NOT NULL DEFAULT 'pending',
  email VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,

  PRIMARY KEY (id),

  CONSTRAINT fk_order_ticket FOREIGN KEY (ticket_id) REFERENCES tickets (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;

DROP TYPE ORDER_STATE;
-- +goose StatementEnd
