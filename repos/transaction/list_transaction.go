package transaction

import (
	"fmt"
	"strings"

	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/lib/pq"
)

var listTransactionRepoLogger = logging.GetLogger(
	"transactionRepo",
	"listTransaction",
)

// ListTransaction implements TransactionRepository.
func (t *transactionRepositoryImpl) ListTransaction(
	customerId, createdAtSort string,
	limit int64,
	offset int64,
) ([]Transaction, error) {
	if createdAtSort != "asc" && createdAtSort != "desc" {
		createdAtSort = "desc"
	}

	var result []Transaction

	lol := 0
	getPlaceholder := func() string {
		lol++
		return fmt.Sprintf("$%d", lol)
	}
	withPlaceholder := func(s string) string {
		return fmt.Sprintf(s, getPlaceholder())
	}

	var (
		conditions []string
		values     []interface{}
	)
	addCondition := func(clause string, value interface{}) {
		conditions = append(conditions, withPlaceholder(clause))
		values = append(values, value)
	}

	if customerId != "" {
		addCondition("customer_id = %s", customerId)
	}

	whereClause := ""
	if len(conditions) != 0 {
		whereClause = strings.Join(conditions, " AND ")
	}

	query := fmt.Sprintf(
		`
            SELECT *
            FROM "transaction"
            %s
            ORDER BY created_at %s
            LIMIT %s
            OFFSET %s
        `,
		whereClause,
		createdAtSort,
		getPlaceholder(),
		getPlaceholder(),
	)

	values = append(values, limit, offset)

	var dbRes struct {
		Transaction
		DbProductIDs        pq.StringArray `db:"product_ids"`
		DbProductQuantities pq.Int64Array  `db:"product_quantities"`
	}

	rows, err := t.db.Queryx(query, values...)

	if err != nil {
		listTransactionRepoLogger.Printf(
			"could not execute query: %s",
			err,
		)
	}

	for rows.Next() {
		if err = rows.StructScan(&dbRes); err != nil {
			listTransactionRepoLogger.Printf(
				"could not parse result into struct: %s",
				err,
			)
			return result, err
		}
		res := Transaction{
			TransactionID: dbRes.TransactionID,
			CustomerID:    dbRes.CustomerID,
			Paid:          dbRes.Paid,
			Change:        dbRes.Change,
			CreatedAt:     dbRes.CreatedAt,
			UpdatedAt:     dbRes.UpdatedAt,
		}
		for _, productId := range dbRes.DbProductIDs {
			res.ProductIDs = append(res.ProductIDs, productId)
		}
		for _, quantity := range dbRes.DbProductQuantities {
			res.ProductQuantities = append(res.ProductQuantities, quantity)
		}
		result = append(result, res)
	}

	return result, nil
}
