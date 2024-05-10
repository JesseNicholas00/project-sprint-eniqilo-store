package customer

import "errors"

var ErrPhoneNumberNotFound = errors.New(
	"customerRepository: no such phone number found",
)
