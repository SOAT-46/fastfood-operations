-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders (
  id SERIAL PRIMARY KEY,
  number VARCHAR(255) NOT NULL,
  status VARCHAR(40) NOT NULL CHECK (status IN ('PENDING', 'RECEIVED', 'PREPARATION', 'READY', 'DELIVERED', 'CANCELLED')),
  received_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE,
  payment_id VARCHAR(255) NOT NULL,
  user_id VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
