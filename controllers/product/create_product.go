package product

import (
	"net/http"

	"github.com/JesseNicholas00/EniqiloStore/services/product"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/JesseNicholas00/EniqiloStore/utils/request"
	"github.com/labstack/echo/v4"
)

var createProductBindLogger = logging.GetLogger(
	"productController",
	"createProduct",
	"bind",
)
var createProductProcessLogger = logging.GetLogger(
	"productController",
	"createProduct",
	"process",
)

func (ctrl *productController) createProduct(c echo.Context) error {
	var req product.CreateProductReq
	if err := request.BindAndValidate(c, &req, createProductBindLogger); err != nil {
		return err
	}

	var res product.CreateProductRes
	if err := ctrl.service.CreateProduct(req, &res); err != nil {
		createProductProcessLogger.Printf(
			"error while processing request: %s", err,
		)

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    res,
	})
}
