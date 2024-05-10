package product

import "github.com/JesseNicholas00/EniqiloStore/utils/logging"

var updateProductLogger = logging.GetLogger(
	"productRepo",
	"update",
)

func (repo *productRepositoryImpl) UpdateProduct(
	product Product,
) (res ProductUpdateResult, err error) {
	query := `
		UPDATE products
		SET 
			product_name = :product_id,
			product_sku = :product_sku,
			product_category = :product_category,
			product_image_url = :product_image_url,
			product_stock = :product_stock,
			product_notes = :product_notes,
			product_price = :product_price,
			product_location = :product_location,
			product_is_available = :product_is_available
		WHERE
			product_id = :product_id
	`
	execRes, err := repo.db.NamedExec(query, product)
	if err != nil {
		updateProductLogger.Printf("could not execute query: %s", err)
		return
	}

	execRowCnt, err := execRes.RowsAffected()
	if err != nil {
		updateProductLogger.Printf("could not execute query: %s", err)
		return
	}

	res.RowsAffected = int(execRowCnt)
	return
}
