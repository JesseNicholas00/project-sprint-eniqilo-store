package customer

import "github.com/JesseNicholas00/EniqiloStore/repos/customer"

type customerServiceImpl struct {
	repo customer.CustomerRepository
}

func NewCustomerService(
	repo customer.CustomerRepository,
) CustomerService {
	return &customerServiceImpl{repo: repo}
}
