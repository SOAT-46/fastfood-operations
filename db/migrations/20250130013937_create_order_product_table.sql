-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS order_product (
  id SERIAL PRIMARY KEY,
  order_id INT NOT NULL,
  product_id INT NOT NULL,
  quantity INT NOT NULL,
  CONSTRAINT fk_order_product_orders FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_product;
-- +goose StatementEnd
