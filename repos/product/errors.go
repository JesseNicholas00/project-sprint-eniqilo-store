package product

import "errors"

var (
	ErrIdNotFound = errors.New("productRepo: id not found")
)
