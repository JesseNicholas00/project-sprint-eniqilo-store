package customer

type CreateCustomerReq struct {
	Name        string `json:"name"        validate:"required,min=5,max=50"`
	PhoneNumber string `json:"phoneNumber" validate:"required,phoneNumber"`
}

type CreateCustomerRes struct {
	UserID      string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
}

type ListCustomerReq struct {
	Name        string `query:"name"`
	PhoneNumber string `query:"phoneNumber"`
}

type ListCustomerRes struct {
	Data []CustomerDTO `json:"data"`
}

type CustomerDTO struct {
	UserId      string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
}
