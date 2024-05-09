package auth

type RegisterStaffReq struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,phoneNumber"`
	Name        string `json:"name"        validate:"required,min=5,max=50"`
	Password    string `json:"password"    validate:"required,min=5,max=15"`
}

type RegisterStaffRes struct {
	UserId      string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

type LoginStaffReq struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,phoneNumber"`
	Password    string `json:"password"    validate:"required,min=5,max=15"`
}

type LoginStaffRes struct {
	UserId      string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
