package auth

import (
	"testing"

	"github.com/JesseNicholas00/EniqiloStore/controllers/auth/mocks"
	"github.com/golang/mock/gomock"
)

//go:generate mockgen -destination mocks/mock_service.go -package mocks github.com/JesseNicholas00/EniqiloStore/services/auth AuthService

func NewControllerWithMockedService(
	t *testing.T,
) (
	mockCtrl *gomock.Controller,
	controller *authController,
	mockedService *mocks.MockAuthService,
) {
	mockCtrl = gomock.NewController(t)
	mockedService = mocks.NewMockAuthService(mockCtrl)
	controller = NewAuthController(mockedService).(*authController)
	return
}
