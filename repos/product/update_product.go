package product

import (
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
)

var updateProductRepoLogger = logging.GetLogger(
	"productRepo",
	"updateProduct",
)

func (repo *productRepositoryImpl) UpdateProduct(product Product) (ProductUpdateResult, error) {
	updateQuery := `
		UPDATE products
		SET product_name = :product_name,
			product_sku = :product_sku,
			product_category = :product_category,
			product_image_url = :product_image_url,
			product_stock = :product_stock,
			product_notes = :product_notes,
			product_price = :product_price,
			product_location = :product_location,
			product_is_available = :product_is_available
		WHERE product_id = :product_id
	`

	results, err := repo.db.NamedExec(updateQuery, product)
	res := ProductUpdateResult{}
	if err != nil {
		updateProductRepoLogger.Printf(
			"error while updateProduct() caused by: %s",
			err,
		)
		return res, err
	}

	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return res, err
	}

	res.RowsAffected = int(rowsAffected)

	return res, nil
}
