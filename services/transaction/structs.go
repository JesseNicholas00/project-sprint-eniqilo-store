package transaction

type ListTransactionReq struct {
	CustomerId string `query:"customerId"`
	Limit      int64  `query:"limit"`
	Offset     int64  `query:"offset"`
	CreatedAt  string `query:"createdAt" validate:"oneof=asc desc"`
}

type ListTransactionRes struct {
	Data []TransactionDTO `json:"data"`
}

type TransactionDTO struct {
	TransactionId  string                  `json:"transactionId"`
	CustomerId     string                  `json:"customerId"`
	ProductDetails []TransactionProductDTO `json:"productDetails"`
	Paid           int64                   `json:"paid"`
	Change         int64                   `json:"change"`
	CreatedAt      string                  `json:"createdAt"`
}

type TransactionProductDTO struct {
	ProductId string `json:"productId"`
	Quantity  int64  `json:"quantity"`
}
