package dummy

import (
	"net/http"

	svc "github.com/JesseNicholas00/EniqiloStore/services/dummy"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/JesseNicholas00/EniqiloStore/utils/request"
	"github.com/labstack/echo/v4"
)

var getDummyBindLogger = logging.GetLogger(
	"dummyController",
	"getDummy",
	"bind",
)
var getDummyProcessLogger = logging.GetLogger(
	"dummyController",
	"getDummy",
	"process",
)

func (ctrl *dummyController) getDummy(c echo.Context) error {
	var req svc.GetDummyReq
	if err := request.BindAndValidate(c, &req, getDummyBindLogger); err != nil {
		return err
	}

	var res svc.GetDummyRes
	if err := ctrl.service.GetDummy(req, &res); err != nil {
		getDummyProcessLogger.Printf(
			"error while processing request: %s", err,
		)

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, res)
}
