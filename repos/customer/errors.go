package customer

import "errors"

var (
	ErrPhoneNumberNotFound = errors.New(
		"customerRepository: no such phone number found",
	)
	ErrIdNotFound = errors.New("customerRepository: no such id found")
)
