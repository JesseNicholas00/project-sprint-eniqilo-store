package customer

import (
	"github.com/JesseNicholas00/EniqiloStore/repos/customer"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
)

var listCustomerServiceLogger = logging.GetLogger(
	"listCustomerService",
	"listCustomer",
)

func (svc *customerServiceImpl) ListCustomer(
	req ListCustomerReq,
	res *ListCustomerRes,
) error {
	customerRes, err := svc.repo.ListCustomer(req.Name, req.PhoneNumber)
	if err != nil {
		listCustomerServiceLogger.Printf(
			"error while ListCustomer() caused by: %s",
			err,
		)
		return err
	}
	*res = ListCustomerRes{
		Data: toDTO(customerRes),
	}
	return nil
}

func toDTO(customers []customer.Customer) []CustomerDTO {
	res := []CustomerDTO{}
	for _, customer := range customers {
		res = append(res, CustomerDTO{
			UserId:      customer.CustomerID,
			PhoneNumber: customer.PhoneNumber,
			Name:        customer.Name,
		})
	}

	return res
}
