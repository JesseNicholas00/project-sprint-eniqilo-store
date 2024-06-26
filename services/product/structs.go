package product

type CreateProductReq struct {
	Name      string `json:"name"        validate:"required,min=1,max=30"`
	SKU       string `json:"sku"         validate:"required,min=1,max=30"`
	Category  string `json:"category"    validate:"required,oneof=Clothing Accessories Footwear Beverages"`
	ImageUrl  string `json:"imageUrl"    validate:"required,url,imageExtension"`
	Notes     string `json:"notes"       validate:"required,min=1,max=200"`
	Price     int    `json:"price"       validate:"required,min=1"`
	Stock     int    `json:"stock"       validate:"required,min=0,max=100000"`
	Location  string `json:"location"    validate:"required,min=1,max=200"`
	Available *bool  `json:"isAvailable" validate:"required"`
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
	ID        string `param:"id" validate:"required"`
	Name      string `           validate:"required,min=1,max=30"                                  json:"name"`
	SKU       string `           validate:"required,min=1,max=30"                                  json:"sku"`
	Category  string `           validate:"required,oneof=Clothing Accessories Footwear Beverages" json:"category"`
	ImageUrl  string `           validate:"required,url,imageExtension"                            json:"imageUrl"`
	Notes     string `           validate:"required,min=1,max=200"                                 json:"notes"`
	Price     int    `           validate:"required,min=1"                                         json:"price"`
	Stock     int    `           validate:"required,min=0,max=100000"                              json:"stock"`
	Location  string `           validate:"required,min=1,max=200"                                 json:"location"`
	Available *bool  `           validate:"required"                                               json:"isAvailable"`
}

type UpdateProductRes struct {
	Message      string `json:"message"`
	RowsAffected int    `json:"rows_affected"`
}

type GetProductsReq struct {
	Id             string `query:"id"`
	Name           string `query:"name"`
	AvailableInput string `query:"isAvailable"`
	Available      *bool
	Category       string `query:"category"`
	SKU            string `query:"sku"`
	PriceSort      string `query:"price"`
	InStockInput   string `query:"inStock"`
	InStock        *bool
	CreatedAt      string `query:"createdAt"`
	Limit          int    `query:"limit"`
	Offset         int    `query:"offset"`
}

type GetProductsRes struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	SKU       string `json:"sku"`
	Category  string `json:"category"`
	ImageUrl  string `json:"imageUrl"`
	Stock     int    `json:"stock"`
	Notes     string `json:"notes"`
	Price     int    `json:"price"`
	Location  string `json:"location"`
	Available bool   `json:"isAvailable"`
	CreatedAt string `json:"createdAt"`
}

type GetProductsByCustomerRes struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	SKU       string `json:"sku"`
	Category  string `json:"category"`
	ImageUrl  string `json:"imageUrl"`
	Stock     int    `json:"stock"`
	Price     int    `json:"price"`
	Location  string `json:"location"`
	CreatedAt string `json:"createdAt"`
}

type CheckoutProductReq struct {
	CustomerId     string          `json:"customerId"     validate:"required"`
	ProductDetails []ProductDetail `json:"productDetails" validate:"required,min=1,dive"`
	Paid           int64           `json:"paid"           validate:"required,min=1"`
	Change         *int64          `json:"change"         validate:"required,min=0"`
}

type ProductDetail struct {
	ProductId string `json:"productId" validate:"required"`
	Quantity  int    `json:"quantity"  validate:"required,min=1"`
}

type CheckoutProductRes struct {
	ProductId string `json:"productId"`
	TotalCost int64  `json:"totalCost"`
}
