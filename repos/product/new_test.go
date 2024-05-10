package product_test

import (
	"testing"

	"github.com/JesseNicholas00/EniqiloStore/repos/product"
	"github.com/JesseNicholas00/EniqiloStore/utils/unittesting"
)

func NewWithTestDatabase(t *testing.T) product.ProductRepository {
	db := unittesting.SetupTestDatabase("../../migrations", t)
	return product.NewProductRepository(db)
}
