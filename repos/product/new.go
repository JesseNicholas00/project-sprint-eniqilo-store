package product

import "github.com/jmoiron/sqlx"

type productRepositoryImpl struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepositoryImpl{db: db}
}
