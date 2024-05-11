package customer

import (
	"fmt"
	"strings"

	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
)

var listCustomerRepoLogger = logging.GetLogger(
	"cutomerRepo",
	"listCustomer",
)

func (repo *customerRepositoryImpl) ListCustomer(
	customerName, customerPhoneNumber string,
) ([]Customer, error) {
	var customers []Customer

	position := 1
	var conditions []string
	var parameters []interface{}

	if customerName != "" {
		conditions = append(conditions, fmt.Sprintf("customer_name ILIKE $%d", position))
		parameters = append(parameters, "%"+customerName+"%")
	}
	if customerPhoneNumber != "" {
		conditions = append(conditions, fmt.Sprintf("customer_phone_number ILIKE $%d", position))
		parameters = append(parameters, "+"+customerPhoneNumber+"%")
	}
	conditionalQuery := ""
	if len(conditions) > 0 {
		conditionalQuery = "WHERE " + strings.Join(conditions, " AND ")
	}
	query := "SELECT * FROM customers " + conditionalQuery + " ORDER BY created_at DESC"
	rows, err := repo.db.Queryx(query, parameters...)

	if err != nil {
		listCustomerRepoLogger.Printf("error while listCustomer() caused by: %s", err)
		return []Customer{}, err
	}
	defer rows.Close()
	for rows.Next() {
		customer := Customer{}
		err := rows.StructScan(&customer)
		if err != nil {
			listCustomerRepoLogger.Printf("error while listCustomer() caused by: %s", err)
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}
