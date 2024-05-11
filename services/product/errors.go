package product

import "errors"

var (
	ErrCustomerNotFound  = errors.New("productService: no such customer found")
	ErrProductNotFound   = errors.New("productService: no such product found")
	ErrProductOutOfStock = errors.New(
		"productService: product does not have enough stock",
	)
	ErrProductInactive = errors.New("productService: product is inactive")
	ErrNotEnoughPaid   = errors.New("productService: not enough money paid")
	ErrIncorrectChange = errors.New("productService: change is incorrect")

	ErrFailedToUpdate = errors.New("productService: could not update product")
)
