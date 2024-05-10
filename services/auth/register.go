package auth

import (
	"errors"

	"github.com/JesseNicholas00/EniqiloStore/repos/auth"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var registerLogger = logging.GetLogger("authService", "register")

func (svc *authServiceImpl) RegisterStaff(
	req RegisterStaffReq,
	res *RegisterStaffRes,
) error {
	_, err := svc.repo.FindStaffByPhone(req.PhoneNumber)

	if err == nil {
		// duplicate phone number
		return ErrPhoneNumberAlreadyRegistered
	}

	if !errors.Is(err, auth.ErrPhoneNumberNotFound) {
		// unexpected kind of error
		registerLogger.Printf(
			"could not check for duplicate phone numbers: %s",
			err,
		)
		return err
	}

	cryptedPw, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		svc.bcryptCost,
	)
	if err != nil {
		registerLogger.Printf("could not crypt password: %s", err)
		return err
	}

	repoRes, err := svc.repo.CreateStaff(auth.Staff{
		Id:       uuid.New().String(),
		Name:     req.Name,
		Phone:    req.PhoneNumber,
		Password: string(cryptedPw),
	})
	if err != nil {
		registerLogger.Printf(
			"repo could not create staff: %s",
			err,
		)
		return err
	}

	token, err := svc.generateToken(repoRes)
	if err != nil {
		registerLogger.Printf("could not generate token: %s", err)
		return err
	}

	*res = RegisterStaffRes{
		UserId:      repoRes.Id,
		PhoneNumber: repoRes.Phone,
		Name:        repoRes.Name,
		AccessToken: token,
	}

	return nil
}
