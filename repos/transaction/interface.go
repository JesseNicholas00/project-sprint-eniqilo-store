package transaction

type TransactionRepository interface {
	ListTransaction(
		customerId, createdAtSort string,
		limit int64,
		offset int64,
	) ([]Transaction, error)
	CreateTransaction(transaction Transaction) (Transaction, error)
}
