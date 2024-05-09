package product

type ProductService interface {
	CreateProduct(req CreateProductReq, res *CreateProductRes) error
}
