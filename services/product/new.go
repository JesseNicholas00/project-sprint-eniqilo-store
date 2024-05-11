package product

import (
	"github.com/JesseNicholas00/EniqiloStore/repos/customer"
	"github.com/JesseNicholas00/EniqiloStore/repos/product"
)

type productServiceImpl struct {
	repo     product.ProductRepository
	custRepo customer.CustomerRepository
}

func NewProductService(
	repo product.ProductRepository,
	custRepo customer.CustomerRepository,
) ProductService {
	return &productServiceImpl{repo: repo, custRepo: custRepo}
}
