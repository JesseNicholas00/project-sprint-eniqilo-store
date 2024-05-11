package customer

type CustomerRepository interface {
	CreateCustomer(Customer Customer) (Customer, error)
	FindCustomerByPhoneNumber(phoneNumber string) (Customer, error)
	ListCustomer(customerName, customerPhoneNumber string) ([]Customer, error)
	GetCustomerById(id string) (Customer, error)
}
