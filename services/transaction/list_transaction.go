package transaction

import (
	"encoding/json"

	productDb "github.com/JesseNicholas00/EniqiloStore/repos/product"
	transactionDb "github.com/JesseNicholas00/EniqiloStore/repos/transaction"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
)

var listTransactionServiceLogger = logging.GetLogger(
	"transactionService",
	"listTransaction",
)

// ListTransaction implements TransactionService.
func (t *transactionServiceImpl) ListTransaction(req ListTransactionReq, res *ListTransactionRes) error {

	if req.Limit == 0 {
		req.Limit = 5
	}

	var transactions []transactionDb.Transaction
	transactions, err := t.trxRepo.ListTransaction(req.CustomerId, req.CreatedAt, req.Limit, req.Offset)

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
	listTransactionServiceLogger.Printf(
		"assigned to res: %v",
		res,
	)

	return nil
}

func getTransactionIds(transactions []transactionDb.Transaction) (result []string, err error) {
	for i := range transactions {
		transaction := transactions[i]

		var transProdIds []string
		err = json.Unmarshal([]byte(transaction.ProductIDs), &transProdIds)
		if err != nil {
			listTransactionServiceLogger.Printf(
				"error while getTransactionIds() caused by json: %s",
				err,
			)
			return result, err
		}

		result = append(result, transProdIds...)
	}

	return result, nil
}

func mapToRes(transactions []transactionDb.Transaction, products []productDb.Product, res *ListTransactionRes) error {
	productMap := make(map[string]productDb.Product)

	for i := range products {
		productMap[products[i].ProductID] = products[i]
	}

	var trans []TransactionDTO

	for i := range transactions {

		transaction := transactions[i]

		var transProdIds []string
		err := json.Unmarshal([]byte(transaction.ProductIDs), &transProdIds)
		if err != nil {
			listTransactionServiceLogger.Printf(
				"error while mapToRes() caused by json: %s",
				err,
			)
			return err
		}
		var transProdAmount []int64
		err = json.Unmarshal([]byte(transaction.ProductQuantities), &transProdAmount)
		if err != nil {
			listTransactionServiceLogger.Printf(
				"error while mapToRes() caused by json: %s",
				err,
			)
			return err
		}

		var transProds []TransactionProductDTO
		for i := range transProdIds {
			product := productMap[transProdIds[i]]
			transProds = append(transProds, TransactionProductDTO{
				ProductId: product.ProductID,
				Quantity:  transProdAmount[i],
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
