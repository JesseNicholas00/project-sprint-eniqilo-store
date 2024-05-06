package dummy

import "github.com/jmoiron/sqlx"

type dummyRepositoryImpl struct {
	db *sqlx.DB
}

func NewDummyRepository(db *sqlx.DB) DummyRepository {
	return &dummyRepositoryImpl{db: db}
}
