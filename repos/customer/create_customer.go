package customer

import "github.com/JesseNicholas00/EniqiloStore/utils/logging"

var createCustomerRepoLogger = logging.GetLogger(
	"cutomerRepo",
	"createCustomer",
)

func (repo *customerRepositoryImpl) CreateCustomer(customer Customer) (Customer, error) {
	insertQuery := `INSERT INTO customers (
			customer_id,
			customer_name,
			customer_phone_number)
		VALUES (
			:customer_id,
			:customer_name,
			:customer_phone_number)
		RETURNING
			customer_id,
			customer_name,
			customer_phone_number,
			created_at,
			updated_at`
	rows, err := repo.db.NamedQuery(insertQuery, customer)
	res := Customer{}
	if err != nil {
		createCustomerRepoLogger.Printf(
			"error while createCustomer() caused by: %s",
			err,
		)
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&res)
		if err != nil {
			createCustomerRepoLogger.Printf(
				"error while createCustomer() caused by: %s",
				err,
			)
			return res, err
		}
	}
	return res, nil
}
