package product

import (
	"errors"
	"net/http"

	"github.com/JesseNicholas00/EniqiloStore/services/product"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/JesseNicholas00/EniqiloStore/utils/request"
	"github.com/labstack/echo/v4"
)

var updateProductBindLogger = logging.GetLogger(
	"productController",
	"update",
	"bind",
)

var updateProductProcessLogger = logging.GetLogger(
	"productController",
	"update",
	"process",
)

func (ctrl *productController) updateProduct(ctx echo.Context) error {
	var req product.UpdateProductReq
	if err := request.BindAndValidate(ctx, &req, updateProductBindLogger); err != nil {
		return err
	}

	if err := ctrl.service.UpdateProduct(req, &product.UpdateProductRes{}); err != nil {
		switch {
		case errors.Is(err, product.ErrProductNotFound):
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": "product not found",
			})

		default:
			updateProductProcessLogger.Printf(
				"could not update product: %s",
				err,
			)
			return ctx.JSON(http.StatusInternalServerError, echo.Map{
				"message": "internal server error",
			})
		}
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
