package product

import (
	"fmt"
	"strings"

	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
)

var getProductsRepoLogger = logging.GetLogger(
	"productRepo",
	"getProducts",
)

func (repo *productRepositoryImpl) GetProducts(productFilter ProductFilter) ([]Product, error) {

	var products []Product

	var conditions []string
	var parameters []interface{}

	conditionalQuery := ""
	orderingQuery := ""

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
		conditionalQuery = "WHERE " + strings.Join(conditions, " AND ")
	}

	var orderings []string

	if productFilter.PriceSort != "" {
		orderings = append(orderings, fmt.Sprintf("product_price %s", productFilter.PriceSort))
	}
	if productFilter.CreatedAtSort != "" {
		orderings = append(orderings, fmt.Sprintf("created_at %s", productFilter.CreatedAtSort))
	}
	if len(orderings) > 0 {
		orderingQuery = "ORDER BY " + strings.Join(orderings, ",")
	}

	query := fmt.Sprintf(
		"SELECT * FROM products %s %s LIMIT %d OFFSET %d",
		conditionalQuery,
		orderingQuery,
		productFilter.Limit,
		productFilter.Offset,
	)

	rows, err := repo.db.Queryx(query, parameters...)
	if err != nil {
		getProductsRepoLogger.Printf("error while getProducts() caused by: %s", err)
		return []Product{}, err
	}
	defer rows.Close()
	for rows.Next() {
		product := Product{}
		err := rows.StructScan(&product)
		if err != nil {
			getProductsRepoLogger.Printf("error while getProducts() caused by: %s", err)
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
