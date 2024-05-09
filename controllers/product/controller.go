package product

import (
	"github.com/JesseNicholas00/EniqiloStore/controllers"
	"github.com/JesseNicholas00/EniqiloStore/middlewares"
	"github.com/JesseNicholas00/EniqiloStore/services/product"
	"github.com/labstack/echo/v4"
)

type productController struct {
	service product.ProductService
	authMw  *middlewares.AuthMiddleware
}

func (ctrl *productController) Register(server *echo.Echo) error {
	urlGroup := server.Group("/v1/product")
	urlGroup.Use(ctrl.authMw.Process)
	urlGroup.POST("", ctrl.CreateProduct)
	return nil
}

func NewProductController(
	service product.ProductService,
	authMw *middlewares.AuthMiddleware,
) controllers.Controller {
	return &productController{
		service: service,
		authMw:  authMw,
	}
}
