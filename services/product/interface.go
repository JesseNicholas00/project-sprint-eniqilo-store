package product

type ProductService interface {
	CreateProduct(req CreateProductReq, res *CreateProductRes) error
	DeleteProduct(req DeleteResultReq, res *DeleteResultRes) error
	UpdateProduct(req UpdateProductReq, res *UpdateProductRes) error
}
