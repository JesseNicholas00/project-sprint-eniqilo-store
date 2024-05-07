CREATE TABLE customers (
    customer_id TEXT,
    customer_name TEXT,
    customer_phone_number TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
