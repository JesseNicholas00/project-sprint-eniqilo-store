package customer

type CustomerRepository interface {
	CreateCustomer(Customer Customer) (Customer, error)
}
