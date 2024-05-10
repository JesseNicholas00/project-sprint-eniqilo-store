package auth_test

import (
	"testing"

	"github.com/JesseNicholas00/EniqiloStore/repos/auth"
	"github.com/JesseNicholas00/EniqiloStore/utils/unittesting"
)

func NewWithTestDatabase(t *testing.T) auth.AuthRepository {
	db := unittesting.SetupTestDatabase("../../migrations", t)
	return auth.NewAuthRepository(db)
}
