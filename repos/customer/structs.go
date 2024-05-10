package customer

type Customer struct {
	CustomerID  string `db:"customer_id"`
	Name        string `db:"customer_name"`
	PhoneNumber string `db:"customer_phone_number"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}
