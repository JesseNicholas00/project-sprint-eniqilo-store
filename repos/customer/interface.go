package customer

type CustomerRepository interface {
	CreateCustomer(Customer Customer) (Customer, error)
	ListCustomer(customerName, customerPhoneNumber string) ([]Customer, error)
}
