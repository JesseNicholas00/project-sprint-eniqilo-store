package product

import (
	"github.com/JesseNicholas00/EniqiloStore/controllers"
	"github.com/JesseNicholas00/EniqiloStore/middlewares"
	"github.com/JesseNicholas00/EniqiloStore/services/product"
	"github.com/labstack/echo/v4"
)

var categories = []string{
	"Clothing",
	"Accessories",
	"Footwear",
	"Beverages",
}

type productController struct {
	service product.ProductService
	authMw  middlewares.Middleware
}

func (ctrl *productController) Register(server *echo.Echo) error {
	urlGroup := server.Group("/v1/product")

	urlGroup.GET("/customer", ctrl.getProductsByCustomer)

	urlGroup.GET("", ctrl.getProducts, ctrl.authMw.Process)
	urlGroup.POST("", ctrl.createProduct, ctrl.authMw.Process)
	urlGroup.DELETE("/:id", ctrl.DeleteProduct, ctrl.authMw.Process)
	urlGroup.PUT("/:id", ctrl.UpdateProduct, ctrl.authMw.Process)
	urlGroup.POST("/checkout", ctrl.checkoutProduct, ctrl.authMw.Process)

	return nil
}

func NewProductController(
	service product.ProductService,
	authMw middlewares.Middleware,
) controllers.Controller {
	return &productController{
		service: service,
		authMw:  authMw,
	}
}
