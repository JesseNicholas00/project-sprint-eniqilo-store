package customer

import (
	"github.com/JesseNicholas00/EniqiloStore/controllers"
	"github.com/JesseNicholas00/EniqiloStore/middlewares"
	"github.com/JesseNicholas00/EniqiloStore/services/customer"
	"github.com/labstack/echo/v4"
)

type customerController struct {
	service customer.CustomerService
	authMw  middlewares.Middleware
}

func (ctrl *customerController) Register(server *echo.Echo) error {
	server.POST("/v1/customer/register", ctrl.CreateCustomer, ctrl.authMw.Process)
	server.GET("/v1/customer", ctrl.ListCustomer)
	return nil
}

func NewCustomerController(
	service customer.CustomerService,
	authMw middlewares.Middleware,
) controllers.Controller {
	return &customerController{
		service: service,
		authMw:  authMw,
	}
}
