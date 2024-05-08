package customer

import (
	"net/http"

	"github.com/JesseNicholas00/EniqiloStore/services/customer"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/JesseNicholas00/EniqiloStore/utils/request"
	"github.com/labstack/echo/v4"
)

var createCustomerBindLogger = logging.GetLogger(
	"customerController",
	"createCustomer",
	"bind",
)
var createCustomerProcessLogger = logging.GetLogger(
	"customerController",
	"createCustomer",
	"process",
)

func (ctrl *customerController) CreateCustomer(c echo.Context) error {
	var req customer.CreateCustomerReq
	if err := request.BindAndValidate(c, &req, createCustomerBindLogger); err != nil {
		return err
	}

	var res customer.CreateCustomerRes
	if err := ctrl.service.CreateCustomer(req, &res); err != nil {
		createCustomerProcessLogger.Printf(
			"error while processing request: %s", err,
		)

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, res)
}
