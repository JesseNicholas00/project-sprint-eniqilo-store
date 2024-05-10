package product

import (
	productRepo "github.com/JesseNicholas00/EniqiloStore/repos/product"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
)

var getProductsServiceLogger = logging.GetLogger(
	"productRepo",
	"getProducts",
)

func (svc *productServiceImpl) GetProducts(req GetProductsReq, res *[]GetProductsRes) error {
	productFilter := productRepo.ProductFilter{
		ProductID:     req.Id,
		Name:          req.Name,
		Available:     req.Available,
		InStock:       req.InStock,
		Category:      req.Category,
		SKU:           req.SKU,
		PriceSort:     req.PriceSort,
		CreatedAtSort: req.CreatedAt,
		Limit:         req.Limit,
		Offset:        req.Offset,
	}
	_, err := svc.repo.GetProducts(productFilter)
	if err != nil {
		getProductsServiceLogger.Printf(
			"error while createProduct() caused by: %s",
			err,
		)
		return err
	}
	// *res = &[]GetProductsRes
	return nil
}
