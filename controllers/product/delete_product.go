package product

import (
	"net/http"

	"github.com/JesseNicholas00/EniqiloStore/services/product"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/JesseNicholas00/EniqiloStore/utils/request"
	"github.com/labstack/echo/v4"
)

var deleteProductBindLogger = logging.GetLogger(
	"productController",
	"deleteProduct",
	"bind",
)

var deleteProductProcessLogger = logging.GetLogger(
	"productController",
	"deleteProduct",
	"process",
)

func (ctrl *productController) DeleteProduct(c echo.Context) error {
	var req product.DeleteProductReq
	if err := request.BindAndValidate(c, &req, createProductBindLogger); err != nil {
		return err
	}

	var res product.DeleteProductRes
	if err := ctrl.service.DeleteProduct(req, &res); err != nil {
		createProductProcessLogger.Printf(
			"error while processing request: %s", err,
		)

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}

	if res.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "product not found",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": res.Message,
	})
}
