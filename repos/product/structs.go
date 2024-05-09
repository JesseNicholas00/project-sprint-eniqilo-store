package product

type Product struct {
	ProductID string `db:"product_id"`
	Name      string `db:"product_name"`
	SKU       string `db:"product_sku"`
	Category  string `db:"product_category"`
	ImageUrl  string `db:"product_image_url"`
	Stock     int    `db:"product_stock"`
	Notes     string `db:"product_notes"`
	Price     int    `db:"product_price"`
	Location  string `db:"product_location"`
	Available bool   `db:"product_is_available"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
