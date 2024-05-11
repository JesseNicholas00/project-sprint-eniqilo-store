package product

import (
	"database/sql"
	"errors"

	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
)

var getByIdLogger = logging.GetLogger(
	"productRepo",
	"getById",
)

func (repo *productRepositoryImpl) GetProductById(
	id string,
) (res Product, err error) {
	query := `SELECT * FROM products WHERE product_id = $1`

	err = repo.db.Get(&res, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrIdNotFound
			return
		}

		getByIdLogger.Printf("could not get from db: %s", err)
		return
	}

	return
}
