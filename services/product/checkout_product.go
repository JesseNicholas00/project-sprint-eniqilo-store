package product

import (
	"errors"

	"github.com/JesseNicholas00/EniqiloStore/repos/customer"
	"github.com/JesseNicholas00/EniqiloStore/repos/product"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
)

var checkoutProductLogger = logging.GetLogger(
	"productService",
	"checkoutProduct",
)

func (svc *productServiceImpl) CheckoutProduct(
	req CheckoutProductReq,
	res *CheckoutProductRes,
) error {
	_, err := svc.custRepo.GetCustomerById(req.CustomerId)
	if err != nil {
		if errors.Is(err, customer.ErrIdNotFound) {
			return ErrCustomerNotFound
		}

		checkoutProductLogger.Printf("could not get customer: %s", err)
		return err
	}

	var productIds []string
	for _, detail := range req.ProductDetails {
		productIds = append(productIds, detail.ProductId)
	}

	products, err := svc.repo.ListProductByIds(productIds)
	if err != nil {
		checkoutProductLogger.Printf("could not get products: %s", err)
	}

	productsById := make(map[string]product.Product)
	for _, product := range products {
		productsById[product.ProductID] = product
	}

	totalCost := int64(0)
	for _, detail := range req.ProductDetails {
		curProduct, found := productsById[detail.ProductId]
		if !found {
			res.ProductId = detail.ProductId
			return ErrProductNotFound
		}

		if curProduct.Stock < detail.Quantity {
			res.ProductId = detail.ProductId
			return ErrProductOutOfStock
		}

		if !curProduct.Available {
			res.ProductId = detail.ProductId
			return ErrProductInactive
		}

		totalCost += int64(curProduct.Price) * int64(detail.Quantity)
	}

	res.TotalCost = totalCost
	if req.Paid < totalCost {
		return ErrNotEnoughPaid
	}

	expectedChange := req.Paid - totalCost
	if *req.Change != expectedChange {
		return ErrIncorrectChange
	}

	for _, detail := range req.ProductDetails {
		curProduct := productsById[detail.ProductId]

		curProduct.Stock -= detail.Quantity

		updRes, err := svc.repo.UpdateProduct(curProduct)

		if err != nil {
			checkoutProductLogger.Printf("could not update product: %s", err)
			return ErrFailedToUpdate
		}

		if updRes.RowsAffected == 0 {
			checkoutProductLogger.Printf(
				"update returned 0 rows updated: %s",
				curProduct.ProductID,
			)
			return ErrFailedToUpdate
		}
	}

	return nil
}
