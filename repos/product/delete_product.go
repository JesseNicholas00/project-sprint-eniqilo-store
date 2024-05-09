package product

import (
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
)

var deleteProductRepoLogger = logging.GetLogger(
	"productRepo",
	"deleteProduct",
)

func (repo *productRepositoryImpl) DeleteProduct(id string) (*ProductDeleteResult, error) {
	deleteQuery := `DELETE FROM products WHERE product_id = $1`

	results, err := repo.db.Exec(deleteQuery, id)
	if err != nil {
		deleteProductRepoLogger.Printf(
			"error while deleteProduct() caused by: %s",
			err,
		)
		return nil, err
	}

	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return nil, err
	}

	res := ProductDeleteResult{
		Message:      "success",
		RowsAffected: int(rowsAffected),
	}
	return &res, nil
}
