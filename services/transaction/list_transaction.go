package transaction

import (
	productDb "github.com/JesseNicholas00/EniqiloStore/repos/product"
	transactionDb "github.com/JesseNicholas00/EniqiloStore/repos/transaction"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
)

var listTransactionServiceLogger = logging.GetLogger(
	"transactionService",
	"listTransaction",
)

// ListTransaction implements TransactionService.
func (t *transactionServiceImpl) ListTransaction(
	req ListTransactionReq,
	res *ListTransactionRes,
) error {

	if req.Limit == 0 {
		req.Limit = 5
	}

	var transactions []transactionDb.Transaction
	transactions, err := t.trxRepo.ListTransaction(
		req.CustomerId,
		req.CreatedAt,
		req.Limit,
		req.Offset,
	)

	if err != nil {
		listTransactionServiceLogger.Printf(
			"error while listTransaction() caused by: %s",
			err,
		)
		return err
	}

	prodIds, err := getTransactionIds(transactions)
	if err != nil {
		listTransactionServiceLogger.Printf(
			"error while listTransaction() caused by getTransactionIds(): %s",
			err,
		)
		return err
	}

	var products []productDb.Product
	if len(transactions) != 0 {
		products, err = t.prdRepo.ListProductByIds(prodIds)
		if err != nil {
			listTransactionServiceLogger.Printf(
				"error while listTransaction() caused by ListProductByIds(): %s",
				err,
			)
			return err
		}
	}

	err = mapToRes(transactions, products, res)
	if err != nil {
		listTransactionServiceLogger.Printf(
			"error while listTransaction() caused by getTransactionIds(): %s",
			err,
		)
		return err
	}

	return nil
}

func getTransactionIds(transactions []transactionDb.Transaction) (result []string, err error) {
	for _, transaction := range transactions {
		result = append(result, transaction.ProductIDs...)
	}
	return result, nil
}

func mapToRes(
	transactions []transactionDb.Transaction,
	products []productDb.Product,
	res *ListTransactionRes,
) error {
	productMap := make(map[string]productDb.Product)

	for _, product := range products {
		productMap[product.ProductID] = product
	}

	var trans []TransactionDTO

	for _, transaction := range transactions {

		var transProds []TransactionProductDTO
		for i, productId := range transaction.ProductIDs {
			transProds = append(transProds, TransactionProductDTO{
				ProductId: productId,
				Quantity:  transaction.ProductQuantities[i],
			})
		}

		trans = append(trans, TransactionDTO{
			TransactionId:  transaction.TransactionID,
			CustomerId:     transaction.CustomerID,
			Paid:           transaction.Paid,
			Change:         transaction.Change,
			CreatedAt:      transaction.CreatedAt,
			ProductDetails: transProds,
		})
	}

	*res = ListTransactionRes{
		Data: trans,
	}

	return nil
}
