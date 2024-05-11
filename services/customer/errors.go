package customer

import "errors"

var ErrPhoneNumberAlreadyRegistered = errors.New(
	"customerService: phone number already registered",
)
