package transaction

type TransactionService interface {
	ListTransaction(req ListTransactionReq, res *ListTransactionRes) error
}
