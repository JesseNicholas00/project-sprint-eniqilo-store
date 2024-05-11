package product

type ProductService interface {
	CreateProduct(req CreateProductReq, res *CreateProductRes) error
	DeleteProduct(req DeleteProductReq, res *DeleteProductRes) error
	UpdateProduct(req UpdateProductReq, res *UpdateProductRes) error
	GetProducts(req GetProductsReq, res *[]GetProductsRes) error
	GetProductsByCustomer(
		req GetProductsReq,
		res *[]GetProductsByCustomerRes,
	) error
	CheckoutProduct(req CheckoutProductReq, res *CheckoutProductRes) error
}
