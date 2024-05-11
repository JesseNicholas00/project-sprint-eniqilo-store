package transaction

import (
	"github.com/jmoiron/sqlx"
)

type transactionRepositoryImpl struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &transactionRepositoryImpl{db: db}
}
