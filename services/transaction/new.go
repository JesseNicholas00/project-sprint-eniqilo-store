package transaction

import (
	"github.com/JesseNicholas00/EniqiloStore/repos/product"
	"github.com/JesseNicholas00/EniqiloStore/repos/transaction"
)

type transactionServiceImpl struct {
	trxRepo transaction.TransactionRepository
	prdRepo product.ProductRepository
}

func NewTransactionService(
	trxRepo transaction.TransactionRepository,
	prdRepo product.ProductRepository,
) TransactionService {
	return &transactionServiceImpl{trxRepo: trxRepo, prdRepo: prdRepo}
}
