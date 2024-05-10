package customer

import (
	"github.com/JesseNicholas00/EniqiloStore/repos/customer"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/google/uuid"
)

var createProductServiceLogger = logging.GetLogger(
	"productRepo",
	"createProduct",
)

func (svc *customerServiceImpl) CreateCustomer(
	req CreateCustomerReq,
	res *CreateCustomerRes,
) error {
	_, err := svc.repo.FindCustomerByPhoneNumber(req.PhoneNumber)

	if err == nil {
		// duplicate phone number
		return ErrPhoneNumberAlreadyRegistered
	}
	product := customer.Customer{
		CustomerID:  uuid.New().String(),
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	}
	savedCustomer, err := svc.repo.CreateCustomer(product)
	if err != nil {
		createProductServiceLogger.Printf(
			"error while createProduct() caused by: %s",
			err,
		)
		return err
	}
	*res = CreateCustomerRes{
		UserID:      savedCustomer.CustomerID,
		PhoneNumber: savedCustomer.PhoneNumber,
		Name:        savedCustomer.Name,
	}
	return nil
}
