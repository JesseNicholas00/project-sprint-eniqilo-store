package product

type CreateProductReq struct {
	Name      string `json:"name"        validate:"required,min=1,max=30"`
	SKU       string `json:"sku"         validate:"required,min=1,max=30"`
	Category  string `json:"category"    validate:"required,oneof=Clothing Accessories Footwear Beverages"`
	ImageUrl  string `json:"imageUrl"    validate:"required,url"`
	Notes     string `json:"notes"       validate:"required,min=1,max=200"`
	Price     int    `json:"price"       validate:"required,min=1"`
	Stock     int    `json:"stock"       validate:"required,min=0,max=100000"`
	Location  string `json:"location"    validate:"required,min=1,max=200"`
	Available bool   `json:"isAvailable" validate:"required"`
}

type CreateProductRes struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

type DeleteProductReq struct {
	ID string `param:"id"        validate:"required"`
}

type DeleteProductRes struct {
	Message      string `json:"message"`
	RowsAffected int    `json:"rows_affected"`
}

type UpdateProductReq struct {
	ID        string `param:"id"        validate:"required,min=1,max=30"`
	Name      string `json:"name"        validate:"required,min=1,max=30"`
	SKU       string `json:"sku"         validate:"required,min=1,max=30"`
	Category  string `json:"category"    validate:"required,oneof=Clothing Accessories Footwear Beverages"`
	ImageUrl  string `json:"imageUrl"    validate:"required,url"`
	Notes     string `json:"notes"       validate:"required,min=1,max=200"`
	Price     int    `json:"price"       validate:"required,min=1"`
	Stock     int    `json:"stock"       validate:"required,min=0,max=100000"`
	Location  string `json:"location"    validate:"required,min=1,max=200"`
	Available bool   `json:"isAvailable" validate:"required"`
}

type UpdateProductRes struct {
	Message      string `json:"message"`
	RowsAffected int    `json:"rows_affected"`
}
