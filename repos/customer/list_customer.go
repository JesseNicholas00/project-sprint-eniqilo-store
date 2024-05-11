package customer

import (
	"database/sql"

	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
)

var listCustomerRepoLogger = logging.GetLogger(
	"cutomerRepo",
	"listCustomer",
)

func (repo *customerRepositoryImpl) ListCustomer(customerName, customerPhoneNumber string) ([]Customer, error) {
	listCustomerRepoLogger.Printf(
		"start ListCustomer() with customerName: %s and customerPhoneNumber: %s",
		customerName, customerPhoneNumber,
	)

	res := []Customer{}
	insertQuery := `SELECT *
	FROM customers`

	plhdr := 1
	wheres := []string{}
	args := []string{}
	if customerName != "" {
		wheres = append(wheres, `customer_name ilike '\%$`+string(plhdr)+`\%'`)
		args = append(args, customerName)
		plhdr++
	}
	if customerPhoneNumber != "" {
		wheres = append(wheres, `customer_phone_number like '\%$+string(plhdr)+\%'`)
		args = append(args, customerPhoneNumber)
		plhdr++
	}
	for i := 0; i < len(wheres); i++ {
		if i == 0 {
			insertQuery = insertQuery + `
			WHERE ` + wheres[i]
		} else {
			insertQuery = insertQuery + `
			AND ` + wheres[i]
		}
	}
	err := repo.db.Select(res, insertQuery, args)

	if err != nil && err != sql.ErrNoRows {
		listCustomerRepoLogger.Printf(
			"error while ListCustomer() caused by: %s",
			err,
		)
		return res, err
	}
	return res, nil
}
