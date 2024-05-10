package customer

import "github.com/jmoiron/sqlx"

type customerRepositoryImpl struct {
	db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	return &customerRepositoryImpl{db: db}
}
