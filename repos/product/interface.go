package product

type ProductRepository interface {
	CreateProduct(Product Product) (Product, error)
}
