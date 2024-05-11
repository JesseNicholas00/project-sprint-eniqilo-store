package customer

import (
	"database/sql"
	"errors"

	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
)

var getByIdLogger = logging.GetLogger(
	"customerRepo",
	"getById",
)

func (repo *customerRepositoryImpl) GetCustomerById(
	id string,
) (res Customer, err error) {
	query := `SELECT * FROM customers WHERE customer_id = $1`

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
