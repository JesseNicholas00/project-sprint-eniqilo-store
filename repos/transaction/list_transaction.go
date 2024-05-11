package transaction

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
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

	addCondition("customer_id = %s", customerId)

	query := fmt.Sprintf(
		`
            SELECT *
            FROM "transaction"
            WHERE %s
            ORDER BY created_at %s
            LIMIT %s
            OFFSET %s
        `,
		strings.Join(conditions, " AND "),
		createdAtSort,
		getPlaceholder(),
		getPlaceholder(),
	)

	values = append(values, limit, offset)

	err := t.db.Select(&result, query, values...)
	if err != nil && err != sql.ErrNoRows {
		listTransactionRepoLogger.Printf(
			"error while listTraction() caused by: %s",
			err,
		)
		return result, err
	}

	return result, nil
}
