package customer

import (
	"net/http"

	"github.com/JesseNicholas00/EniqiloStore/services/customer"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/JesseNicholas00/EniqiloStore/utils/request"
	"github.com/labstack/echo/v4"
)

var listCustomerBindLogger = logging.GetLogger(
	"customerController",
	"listCustomer",
	"bind",
)
var listCustomerProcessLogger = logging.GetLogger(
	"customerController",
	"listCustomer",
	"process",
)

func (ctrl *customerController) ListCustomer(c echo.Context) error {
	var req customer.ListCustomerReq
	if err := request.BindAndValidate(c, &req, listCustomerBindLogger); err != nil {
		return err
	}

	var res customer.ListCustomerRes
	if err := ctrl.service.ListCustomer(req, &res); err != nil {
		listCustomerProcessLogger.Printf(
			"error while processing request: %s", err,
		)

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, res)
}
