package customer

type CreateCustomerReq struct {
	Name        string `json:"name"        validate:"required,min=5,max=50"`
	PhoneNumber string `json:"phoneNumber" validate:"required,e164"`
}

type CreateCustomerRes struct {
	UserID      string `json:"user_id"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
}
