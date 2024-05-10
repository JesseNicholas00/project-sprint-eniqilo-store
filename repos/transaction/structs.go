package transaction

type Transaction struct {
	TransactionID     string `db:"transaction_id"`
	CustomerID        string `db:"customer_id"`
	ProductIDs        string `db:"product_ids"`
	ProductQuantities string `db:"product_quantities"`
	Paid              int64  `db:"paid"`
	Change            int64  `db:"change"`
	CreatedAt         string `db:"created_at"`
	UpdatedAt         string `db:"updated_at"`
}
