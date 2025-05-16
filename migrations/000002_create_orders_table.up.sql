CREATE TABLE orders
(
    id           SERIAL PRIMARY KEY,
    user_id      INTEGER     NOT NULL,
    order_number VARCHAR(64) NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at   TIMESTAMPTZ,

    CONSTRAINT orders_order_number_unique UNIQUE (order_number),
    CONSTRAINT orders_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE INDEX idx_orders_user_id ON orders (user_id);
