package customer

import (
	"github.com/JesseNicholas00/EniqiloStore/controllers"
	"github.com/JesseNicholas00/EniqiloStore/services/customer"
	"github.com/labstack/echo/v4"
)

type customerController struct {
	service customer.CustomerService
}

func (ctrl *customerController) Register(server *echo.Echo) error {
	server.POST("/v1/customer/register", ctrl.CreateCustomer)
	return nil
}

func NewCustomerController(service customer.CustomerService) controllers.Controller {
	return &customerController{
		service: service,
	}
}
