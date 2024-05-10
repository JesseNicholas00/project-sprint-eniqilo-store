package auth

import (
	"errors"

	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/golang-jwt/jwt/v4"
)

var sessionLogger = logging.GetLogger("authService", "session")

func (svc *authServiceImpl) GetSessionFromToken(
	req GetSessionFromTokenReq,
	res *GetSessionFromTokenRes,
) error {
	claims := jwtClaims{}
	_, err := jwt.ParseWithClaims(
		req.AccessToken,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			return svc.jwtSecret, nil
		},
	)

	if err != nil {
		switch {

		case errors.Is(err, jwt.ErrTokenMalformed) ||
			errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return ErrTokenInvalid

		case errors.Is(err, jwt.ErrTokenExpired) ||
			errors.Is(err, jwt.ErrTokenNotValidYet):
			return ErrTokenExpired

		default:
			sessionLogger.Printf("could not parse access token: %s", err)
			return err
		}
	}

	*res = GetSessionFromTokenRes{
		UserId:      claims.Data.UserId,
		Name:        claims.Data.Name,
		PhoneNumber: claims.Data.PhoneNumber,
	}

	return nil
}
