package product

type ProductRepository interface {
	CreateProduct(Product Product) (Product, error)
	DeleteProduct(id string) (ProductDeleteResult, error)
	UpdateProduct(product Product) (ProductUpdateResult, error)
}
