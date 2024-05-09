package auth

import (
	"net/http"

	"github.com/JesseNicholas00/EniqiloStore/services/auth"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/JesseNicholas00/EniqiloStore/utils/request"
	"github.com/labstack/echo/v4"
)

var registerBindLogger = logging.GetLogger(
	"authController",
	"registerStaff",
	"bind",
)

var registerProcessLogger = logging.GetLogger(
	"authController",
	"registerStaff",
	"process",
)

func (ctrl *authController) registerStaff(c echo.Context) error {
	var req auth.RegisterStaffReq
	if err := request.BindAndValidate(c, &req, registerBindLogger); err != nil {
		return err
	}

	var res auth.RegisterStaffRes
	if err := ctrl.service.RegisterStaff(req, &res); err != nil {
		registerProcessLogger.Printf("error while processing request: %s", err)

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "User registered successfully",
		"data":    res,
	})
}
