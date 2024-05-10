package product

import (
	"fmt"
	"strings"

	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
)

var listProductByIdsRepoLogger = logging.GetLogger(
	"productRepo",
	"listProductByIds",
)

// ListProductByIds implements ProductRepository.
func (repo *productRepositoryImpl) ListProductByIds(id []string) ([]Product, error) {
	var res []Product

	if len(id) == 0 {
		return res, nil
	}

	var (
		placeholders []string
		values       []interface{}
	)
	for i, curId := range id {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
		values = append(values, curId)
	}
	sql := fmt.Sprintf("SELECT * FROM products WHERE product_id in (%s)", strings.Join(placeholders, ","))

	err := repo.db.Select(&res, sql, values...)
	if err != nil {
		listProductByIdsRepoLogger.Printf(
			"error while deleteProduct() caused by: %s",
			err,
		)
		return res, err
	}
	return res, nil
}
