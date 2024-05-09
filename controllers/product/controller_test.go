package product

import (
	"testing"

	"github.com/JesseNicholas00/EniqiloStore/controllers/product/mocks"
	"github.com/golang/mock/gomock"
)

//go:generate mockgen -destination mocks/mock_service.go -package mocks github.com/JesseNicholas00/EniqiloStore/services/product ProductService

func NewControllerWithMockedService(
	t *testing.T,
) (
	mockCtrl *gomock.Controller,
	controller *productController,
	mockedService *mocks.MockProductService,
) {
	mockCtrl = gomock.NewController(t)
	mockedService = mocks.NewMockProductService(mockCtrl)
	controller = NewProductController(mockedService).(*productController)
	return
}
