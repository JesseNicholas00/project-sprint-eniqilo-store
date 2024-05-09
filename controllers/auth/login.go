package auth

import (
	"net/http"

	"github.com/JesseNicholas00/EniqiloStore/services/auth"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/JesseNicholas00/EniqiloStore/utils/request"
	"github.com/labstack/echo/v4"
)

var loginBindLogger = logging.GetLogger(
	"authController",
	"loginStaff",
	"bind",
)

var loginProcessLogger = logging.GetLogger(
	"authController",
	"loginStaff",
	"process",
)

func (ctrl *authController) loginStaff(c echo.Context) error {
	var req auth.LoginStaffReq
	if err := request.BindAndValidate(c, &req, loginBindLogger); err != nil {
		return err
	}

	var res auth.LoginStaffRes
	if err := ctrl.service.LoginStaff(req, &res); err != nil {
		loginProcessLogger.Printf("error while processing request: %s", err)

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User logged in successfully",
		"data":    res,
	})
}
