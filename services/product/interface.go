package product

type ProductService interface {
	CreateProduct(req CreateProductReq, res *CreateProductRes) error
	DeleteProduct(req DeleteProductReq, res *DeleteProductRes) error
	GetProducts(req GetProductsReq, res *[]GetProductsRes) error
}
