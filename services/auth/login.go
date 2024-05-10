package auth

import (
	"github.com/JesseNicholas00/EniqiloStore/repos/auth"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

var loginLogger = logging.GetLogger("authService", "login")

func (svc *authServiceImpl) LoginStaff(
	req LoginStaffReq,
	res *LoginStaffRes,
) error {
	staff, err := svc.repo.FindStaffByPhone(req.PhoneNumber)

	if err != nil {
		if errors.Is(err, auth.ErrPhoneNumberNotFound) {
			return ErrUserNotFound
		}

		loginLogger.Printf("could not find staff by phone: %s", err)
		return err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(staff.Password),
		[]byte(req.Password),
	)
	if err != nil {
		return ErrInvalidCredentials
	}

	token, err := svc.generateToken(staff)
	if err != nil {
		registerLogger.Printf("could not generate token: %s", err)
		return err
	}

	*res = LoginStaffRes{
		UserId:      staff.Id,
		PhoneNumber: staff.Phone,
		Name:        staff.Name,
		AccessToken: token,
	}

	return nil
}
