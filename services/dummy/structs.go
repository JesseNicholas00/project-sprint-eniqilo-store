package dummy

type GetDummyReq struct {
	Id   string `json:"id"   param:"id" validate:"required,oneof=gyatt skibidi amogus"`
	Test int    `json:"test"            validate:"required"                            query:"test"`
}

type GetDummyRes struct {
	Rizz int `json:"rizz"`
}
