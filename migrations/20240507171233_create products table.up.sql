CREATE TABLE products (
    product_id TEXT,
    product_name TEXT,
    product_sku TEXT,
    product_category TEXT,
    product_image_url TEXT,
    product_stock INT,
    product_notes TEXT,
    product_price INT,
    product_location TEXT,
    product_is_available BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
