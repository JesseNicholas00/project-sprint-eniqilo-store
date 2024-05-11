package product

import (
	"github.com/JesseNicholas00/EniqiloStore/repos/customer"
	"github.com/JesseNicholas00/EniqiloStore/repos/product"
	"github.com/JesseNicholas00/EniqiloStore/repos/transaction"
)

type productServiceImpl struct {
	repo     product.ProductRepository
	custRepo customer.CustomerRepository
	trxRepo  transaction.TransactionRepository
}

func NewProductService(
	repo product.ProductRepository,
	custRepo customer.CustomerRepository,
	trxRepo transaction.TransactionRepository,
) ProductService {
	return &productServiceImpl{repo: repo,
		custRepo: custRepo,
		trxRepo:  trxRepo,
	}
}
