package product

import (
	"github.com/JesseNicholas00/EniqiloStore/controllers"
	"github.com/JesseNicholas00/EniqiloStore/services/product"
	"github.com/labstack/echo/v4"
)

type productController struct {
	service product.ProductService
}

func (ctrl *productController) Register(server *echo.Echo) error {
	server.POST("/v1/product", ctrl.CreateProduct)
	return nil
}

func NewProductController(service product.ProductService) controllers.Controller {
	return &productController{
		service: service,
	}
}
