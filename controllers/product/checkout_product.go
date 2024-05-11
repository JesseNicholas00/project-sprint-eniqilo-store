package product

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/JesseNicholas00/EniqiloStore/services/product"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/JesseNicholas00/EniqiloStore/utils/request"
	"github.com/labstack/echo/v4"
)

var (
	checkoutBindLogger = logging.GetLogger(
		"productController",
		"checkout",
		"bind",
	)
	checkoutProcessLogger = logging.GetLogger(
		"productController",
		"checkout",
		"process",
	)
)

func (ctrl *productController) checkoutProduct(ctx echo.Context) error {
	var req product.CheckoutProductReq
	if err := request.BindAndValidate(ctx, &req, checkoutBindLogger); err != nil {
		return err
	}

	var res product.CheckoutProductRes
	if err := ctrl.service.CheckoutProduct(req, &res); err != nil {
		switch {
		case errors.Is(err, product.ErrCustomerNotFound):
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": "customer not found",
			})

		case errors.Is(err, product.ErrProductNotFound):
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": fmt.Sprintf(
					"product not found: %s",
					res.ProductId,
				),
			})

		case errors.Is(err, product.ErrProductOutOfStock):
			return ctx.JSON(http.StatusBadRequest, echo.Map{
				"message": fmt.Sprintf(
					"product out of stock: %s",
					res.ProductId,
				),
			})

		case errors.Is(err, product.ErrProductInactive):
			return ctx.JSON(http.StatusBadRequest, echo.Map{
				"message": fmt.Sprintf(
					"product inactive: %s",
					res.ProductId,
				),
			})

		case errors.Is(err, product.ErrNotEnoughPaid):
			return ctx.JSON(http.StatusBadRequest, echo.Map{
				"message": fmt.Sprintf(
					"not enough paid: expected %d",
					res.TotalCost,
				),
			})

		case errors.Is(err, product.ErrIncorrectChange):
			return ctx.JSON(http.StatusBadRequest, echo.Map{
				"message": fmt.Sprintf(
					"incorrect change: expected %d",
					req.Paid-res.TotalCost,
				),
			})

		default:
			checkoutProcessLogger.Printf("could not checkout: %s", err)
			return ctx.JSON(http.StatusInternalServerError, echo.Map{
				"message": "internal server error",
			})
		}
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
