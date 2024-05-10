package product

import "errors"

var (
	ErrProductNotFound = errors.New("productService: no such product found")
)
