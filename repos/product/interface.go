package product

type ProductRepository interface {
	CreateProduct(Product Product) (Product, error)
	DeleteProduct(id string) (ProductDeleteResult, error)
	UpdateProduct(Product Product) (ProductUpdateResult, error)
	GetProducts(productFilter ProductFilter) ([]Product, error)
	ListProductByIds(id []string) ([]Product, error)
}
