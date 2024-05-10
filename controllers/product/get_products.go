package product

import (
	"net/http"
	"strconv"

	"github.com/JesseNicholas00/EniqiloStore/services/product"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/labstack/echo/v4"
)

var getProductsProcessLogger = logging.GetLogger(
	"productController",
	"getProducts",
	"process",
)

var categories = []string{
	"Clothing",
	"Accessories",
	"Footwear",
	"Beverages",
}

func (ctrl *productController) getProducts(c echo.Context) error {
	var req product.GetProductsReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "invalid request",
		})
	}
	if req.Limit == 0 {
		req.Limit = 5
	}
	available, err := strconv.ParseBool(req.AvailableInput)
	if err == nil {
		req.Available = &available
	}
	selectedCategory := ""
	for _, category := range categories {
		if req.Category == category {
			selectedCategory = category
			break
		}
	}
	req.Category = selectedCategory
	if req.PriceSort != "asc" && req.PriceSort != "desc" {
		req.PriceSort = ""
	}
	inStock, err := strconv.ParseBool(req.InStockInput)
	if err == nil {
		req.InStock = &inStock
	}
	if req.CreatedAt != "asc" && req.CreatedAt != "desc" {
		req.CreatedAt = ""
	}

	res := []product.GetProductsRes{}
	if err := ctrl.service.GetProducts(req, &res); err != nil {
		getProductsProcessLogger.Printf(
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
