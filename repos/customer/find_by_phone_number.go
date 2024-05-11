package customer

import "github.com/JesseNicholas00/EniqiloStore/utils/logging"

var findCustomerByPhoneNumber = logging.GetLogger(
	"customerRepository",
	"findCustomerByPhone",
)

func (repo *customerRepositoryImpl) FindCustomerByPhoneNumber(
	phone string,
) (res Customer, err error) {
	query := `
		SELECT
			*
		FROM
			customers	
		WHERE
			customer_phone_number = :phone_number
	`
	rows, err := repo.db.NamedQuery(query, map[string]interface{}{
		"phone_number": phone,
	})
	if err != nil {
		findCustomerByPhoneNumber.Printf(
			"error executing query: %s",
			err,
		)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&res)
		if err != nil {
			findCustomerByPhoneNumber.Printf(
				"error reading result: %s",
				err,
			)
		}
	}

	if res.CustomerID == "" {
		err = ErrPhoneNumberNotFound
		return
	}

	return
}
