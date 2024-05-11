package transaction

import (
	"net/http"

	"github.com/JesseNicholas00/EniqiloStore/services/transaction"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/JesseNicholas00/EniqiloStore/utils/request"
	"github.com/labstack/echo/v4"
)

var listTransactionBindLogger = logging.GetLogger(
	"transactionController",
	"listTransaction",
	"bind",
)

var listTransactionProcessLogger = logging.GetLogger(
	"transactionController",
	"listTransaction",
	"process",
)

func (t *transactionController) ListTransaction(c echo.Context) error {
	var req transaction.ListTransactionReq
	if err := request.BindAndValidate(c, &req, listTransactionBindLogger); err != nil {
		return err
	}

	var res transaction.ListTransactionRes
	if err := t.service.ListTransaction(req, &res); err != nil {
		listTransactionProcessLogger.Printf(
			"error while processing request: %s", err,
		)

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    res.Data,
	})
}
