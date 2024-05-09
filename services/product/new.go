package product

import "github.com/JesseNicholas00/EniqiloStore/repos/product"

type productServiceImpl struct {
	repo product.ProductRepository
}

func NewProductService(
	repo product.ProductRepository,
) ProductService {
	return &productServiceImpl{repo: repo}
}
