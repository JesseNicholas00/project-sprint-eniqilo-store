package product

import (
	"testing"

	"github.com/JesseNicholas00/EniqiloStore/services/product/mocks"
	gomock "github.com/golang/mock/gomock"
)

//go:generate mockgen -destination mocks/mock_repo.go -package mocks github.com/JesseNicholas00/EniqiloStore/repos/product ProductRepository
//go:generate mockgen -destination mocks/mock_cust_repo.go -package mocks github.com/JesseNicholas00/EniqiloStore/repos/customer CustomerRepository
//go:generate mockgen -destination mocks/mock_trx_repo.go -package mocks github.com/JesseNicholas00/EniqiloStore/repos/transaction TransactionRepository

func NewWithMockedRepo(
	t *testing.T,
) (
	mockCtrl *gomock.Controller,
	service *productServiceImpl,
	mockedRepo *mocks.MockProductRepository,
) {
	mockCtrl = gomock.NewController(t)
	mockedRepo = mocks.NewMockProductRepository(mockCtrl)
	service = NewProductService(mockedRepo, mocks.NewMockCustomerRepository(mockCtrl), mocks.NewMockTransactionRepository(mockCtrl)).(*productServiceImpl)
	return
}

func NewWithCompleteMockedRepos(
	t *testing.T,
) (
	mockCtrl *gomock.Controller,
	service *productServiceImpl,
	mockedRepo *mocks.MockProductRepository,
	mockedCustomerRepo *mocks.MockCustomerRepository,
	mockedTransactionRepo *mocks.MockTransactionRepository,
) {
	mockCtrl = gomock.NewController(t)
	mockedRepo = mocks.NewMockProductRepository(mockCtrl)
	mockedCustomerRepo = mocks.NewMockCustomerRepository(mockCtrl)
	mockedTransactionRepo = mocks.NewMockTransactionRepository(mockCtrl)
	service = NewProductService(
		mockedRepo,
		mockedCustomerRepo,
		mockedTransactionRepo,
	).(*productServiceImpl)
	return
}
