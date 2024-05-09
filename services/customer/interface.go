package customer

type CustomerService interface {
	CreateCustomer(req CreateCustomerReq, res *CreateCustomerRes) error
}
