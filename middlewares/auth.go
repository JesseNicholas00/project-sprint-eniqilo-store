package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/JesseNicholas00/EniqiloStore/services/auth"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/labstack/echo/v4"
)

type authMiddleware struct {
	service auth.AuthService
}

var authMwLogger = logging.GetLogger("authMiddleware")

func (mw *authMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := struct {
			Bearer string `header:"Authorization"`
		}{}

		if err := c.Bind(&header); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "invalid request",
			})
		}

		if header.Bearer == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "missing bearer token",
			})
		}

		splitByBearer := strings.Split(header.Bearer, "Bearer ")
		if len(splitByBearer) != 2 {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "malformed bearer token",
			})
		}

		token := splitByBearer[1]
		req := auth.GetSessionFromTokenReq{
			AccessToken: token,
		}
		var res auth.GetSessionFromTokenRes
		if err := mw.service.GetSessionFromToken(req, &res); err != nil {
			switch {
			case errors.Is(err, auth.ErrTokenInvalid):
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": "malformed bearer token",
				})

			case errors.Is(err, auth.ErrTokenExpired):
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "expired bearer token",
				})

			default:
				authMwLogger.Printf(
					"could not get session from token: %s",
					err,
				)
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "internal server error",
				})
			}
		}

		c.Set("session", res)

		return next(c)
	}
}

func NewAuthMiddleware(service auth.AuthService) *authMiddleware {
	return &authMiddleware{
		service: service,
	}
}
