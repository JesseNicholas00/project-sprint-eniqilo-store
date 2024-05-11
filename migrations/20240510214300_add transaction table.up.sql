CREATE TABLE "transaction"
(
    transaction_id TEXT,
    customer_id TEXT,
    product_ids TEXT[],
    product_quantities BIGINT[],
    paid BIGINT,
    change BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
