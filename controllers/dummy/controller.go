package dummy

import (
	"github.com/JesseNicholas00/EniqiloStore/controllers"
	svc "github.com/JesseNicholas00/EniqiloStore/services/dummy"
	"github.com/labstack/echo/v4"
)

type dummyController struct {
	service svc.DummyService
}

func (ctrl *dummyController) Register(server *echo.Echo) error {
	server.GET("/dummy/:id", ctrl.getDummy)
	return nil
}

func NewDummyController(service svc.DummyService) controllers.Controller {
	return &dummyController{
		service: service,
	}
}
