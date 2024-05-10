package product

import (
	"testing"

	"github.com/JesseNicholas00/EniqiloStore/services/product/mocks"
	gomock "github.com/golang/mock/gomock"
)

//go:generate mockgen -destination mocks/mock_repo.go -package mocks github.com/JesseNicholas00/EniqiloStore/repos/product ProductRepository

func NewWithMockedRepo(
	t *testing.T,
) (
	mockCtrl *gomock.Controller,
	service *productServiceImpl,
	mockedRepo *mocks.MockProductRepository,
) {
	mockCtrl = gomock.NewController(t)
	mockedRepo = mocks.NewMockProductRepository(mockCtrl)
	service = NewProductService(mockedRepo).(*productServiceImpl)
	return
}
