CREATE TABLE "transaction"
(
    transaction_id TEXT,
    customer_id TEXT,
    product_ids TEXT,
    product_quantities TEXT,
    paid BIGINT,
    change BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_customer
      FOREIGN KEY(customer_id) 
        REFERENCES customers(customer_id),
    CONSTRAINT fk_product
      FOREIGN KEY(product_id) 
        REFERENCES product(product_id)
);
