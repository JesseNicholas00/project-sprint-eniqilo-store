package transaction

import (
	"github.com/JesseNicholas00/EniqiloStore/controllers"
	"github.com/JesseNicholas00/EniqiloStore/services/transaction"
	"github.com/labstack/echo/v4"
)

type transactionController struct {
	service transaction.TransactionService
	// authMw  middlewares.Middleware
}

func (ctrl *transactionController) Register(server *echo.Echo) error {
	urlGroup := server.Group("/v1/product/checkout")
	// urlGroup.Use(ctrl.authMw.Process)
	urlGroup.GET("/history", ctrl.ListTransaction)

	return nil
}

func NewTransactionController(
	service transaction.TransactionService,
	// authMw middlewares.Middleware,
) controllers.Controller {
	return &transactionController{
		service: service,
		// authMw:  authMw,
	}
}
