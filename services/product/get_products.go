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
	products, err := svc.repo.GetProducts(productFilter)
	if err != nil {
		getProductsServiceLogger.Printf(
			"error while getProducts() caused by: %s",
			err,
		)
		return err
	}
	for _, product := range products {
		productRes := GetProductsRes{
			Id:        product.ProductID,
			Name:      product.Name,
			SKU:       product.SKU,
			Category:  product.Category,
			ImageUrl:  product.ImageUrl,
			Stock:     product.Stock,
			Notes:     product.Notes,
			Price:     product.Price,
			Location:  product.Location,
			Available: product.Available,
			CreatedAt: product.CreatedAt,
		}
		*res = append(*res, productRes)
	}
	return nil
}
