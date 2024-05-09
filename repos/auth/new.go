package auth

import "github.com/jmoiron/sqlx"

type authRepostioryImpl struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) AuthRepository {
	return &authRepostioryImpl{db: db}
}
