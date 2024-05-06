package request

import (
	"log"
	"net/http"

	"github.com/KerakTelor86/GoBoiler/utils/validation"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func BindAndValidate[R any](
	ctx echo.Context,
	req *R,
	logger *log.Logger,
) error {
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"error": "bad request format (hint: check parameter types)",
		})
	}

	if err := ctx.Validate(req); err != nil {
		if err, ok := err.(validator.ValidationErrors); ok {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"error": validation.ConvertToErrList(err),
			})
		}

		logger.Printf("error while validating: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"error": "internal server error",
		})
	}

	return nil
}
