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
	ID string `param:"id" validate:"required"`
}

type DeleteProductRes struct {
	Message      string `json:"message"`
	RowsAffected int    `json:"rows_affected"`
}

type UpdateProductReq struct {
	Id        string `param:"id" validate:"required"`
	Name      string `           validate:"required,min=1,max=30"                                  json:"name"`
	SKU       string `           validate:"required,min=1,max=30"                                  json:"sku"`
	Category  string `           validate:"required,oneof=Clothing Accessories Footwear Beverages" json:"category"`
	ImageUrl  string `           validate:"required,url"                                           json:"imageUrl"`
	Notes     string `           validate:"required,min=1,max=200"                                 json:"notes"`
	Price     int    `           validate:"required,min=1"                                         json:"price"`
	Stock     int    `           validate:"required,min=0,max=100000"                              json:"stock"`
	Location  string `           validate:"required,min=1,max=200"                                 json:"location"`
	Available bool   `           validate:"required"                                               json:"isAvailable"`
}

type UpdateProductRes struct {
}
