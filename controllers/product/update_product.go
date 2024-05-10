package product

import (
	"net/http"

	"github.com/JesseNicholas00/EniqiloStore/services/product"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/JesseNicholas00/EniqiloStore/utils/request"
	"github.com/labstack/echo/v4"
)

var updateProductBindLogger = logging.GetLogger(
	"productController",
	"updateProduct",
	"bind",
)

var updateProductProcessLogger = logging.GetLogger(
	"productController",
	"updateProduct",
	"process",
)

func (ctrl *productController) UpdateProduct(c echo.Context) error {
	var req product.UpdateProductReq
	if err := request.BindAndValidate(c, &req, updateProductBindLogger); err != nil {
		return err
	}

	var res product.UpdateProductRes
	if err := ctrl.service.UpdateProduct(req, &res); err != nil {
		updateProductProcessLogger.Printf(
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
