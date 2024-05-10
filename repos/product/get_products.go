package product

import (
	"fmt"
	"strings"

	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
)

var logger = logging.GetLogger(
	"productRepo",
	"getProduct",
)

func (repo *productRepositoryImpl) GetProducts(productFilter ProductFilter) ([]Product, error) {

	var products []Product

	var conditions []string
	var parameters []interface{}

	conditionalQuery := ""

	if productFilter.ProductID != "" {
		conditions = append(conditions, "product_id= ?")
		parameters = append(parameters, productFilter.ProductID)
	}
	if productFilter.Name != "" {
		conditions = append(conditions, "product_name = ILIKE '%?%'")
		parameters = append(parameters, productFilter.Name)
	}
	if productFilter.Category != "" {
		conditions = append(conditions, "product_category= ?")
		parameters = append(parameters, productFilter.Category)
	}
	if productFilter.SKU != "" {
		conditions = append(conditions, "product_sku = ?")
		parameters = append(parameters, productFilter.SKU)
	}
	if productFilter.Available != nil {
		conditions = append(conditions, "product_available = ?")
		parameters = append(parameters, *productFilter.Available)
	}
	if productFilter.InStock != nil {
		conditions = append(conditions, "product_in_stock = ?")
		parameters = append(parameters, *productFilter.InStock)
	}

	position := 1
	for idx, condition := range conditions {
		conditions[idx] = strings.ReplaceAll(condition, "?", fmt.Sprintf("%d", position))
		position++
	}
	if len(conditions) > 0 {
		conditionalQuery = "WHERE" + strings.Join(conditions, " AND ")
	}

	query := fmt.Sprintf("SELECT * FROM products %s", conditionalQuery)

	return []Product{}, nil
}
