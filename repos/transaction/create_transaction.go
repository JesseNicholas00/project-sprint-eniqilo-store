package transaction

import (
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

var createTxLogger = logging.GetLogger(
	"transactionRepository",
	"createTransaction",
)

func (repo *transactionRepositoryImpl) CreateTransaction(
	trx Transaction,
) (res Transaction, err error) {
	query := `
		INSERT INTO "transaction"(
			transaction_id,
			customer_id,
			product_ids,
			product_quantities,
			paid,
			change
		) 
		VALUES (
			:transaction_id,
			:customer_id,
			:product_ids,
			:product_quantities,
			:paid,
			:change
		)
		RETURNING
			transaction_id,
			customer_id,
			product_ids,
			product_quantities,
			paid,
			change,
			created_at,
			updated_at
	`

	rows, err := repo.db.NamedQuery(query, map[string]interface{}{
		"transaction_id":     uuid.New().String(),
		"customer_id":        trx.CustomerID,
		"product_ids":        pq.StringArray(trx.ProductIDs),
		"product_quantities": pq.Int64Array(trx.ProductQuantities),
		"paid":               trx.Paid,
		"change":             trx.Change,
	})

	if err != nil {
		createTxLogger.Printf("could not execute sql: %s", err)
		return
	}
	defer rows.Close()

	var dbRes struct {
		Transaction
		DbProductIDs        pq.StringArray `db:"product_ids"`
		DbProductQuantities pq.Int64Array  `db:"product_quantities"`
	}

	for rows.Next() {
		if err = rows.StructScan(&dbRes); err != nil {
			createTxLogger.Printf(
				"could not parse result into struct: %s",
				err,
			)
			return
		}
	}

	res = Transaction{
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

	return
}
