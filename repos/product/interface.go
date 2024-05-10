package product

type ProductRepository interface {
	CreateProduct(Product Product) (Product, error)
	DeleteProduct(id string) (ProductDeleteResult, error)
	ListProductByIds(id []string) ([]Product, error)
}
