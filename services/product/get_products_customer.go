package product

import (
	productRepo "github.com/JesseNicholas00/EniqiloStore/repos/product"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
)

var getProductsByCustomerServiceLogger = logging.GetLogger(
	"productRepo",
	"getProductsByCustomer",
)

func (svc *productServiceImpl) GetProductsByCustomer(
	req GetProductsReq,
	res *[]GetProductsByCustomerRes,
) error {
	productFilter := productRepo.ProductFilter{
		Name:          req.Name,
		Category:      req.Category,
		SKU:           req.SKU,
		PriceSort:     req.PriceSort,
		CreatedAtSort: "asc",
		Limit:         req.Limit,
		InStock:       req.InStock,
		Offset:        req.Offset,
	}
	available := true
	productFilter.Available = &available
	products, err := svc.repo.GetProducts(productFilter)
	if err != nil {
		getProductsByCustomerServiceLogger.Printf(
			"error while getProductsByCustomer() caused by: %s",
			err,
		)
		return err
	}
	for _, product := range products {
		productRes := GetProductsByCustomerRes{
			Id:        product.ProductID,
			Name:      product.Name,
			SKU:       product.SKU,
			Category:  product.Category,
			ImageUrl:  product.ImageUrl,
			Stock:     product.Stock,
			Price:     product.Price,
			Location:  product.Location,
			CreatedAt: product.CreatedAt,
		}
		*res = append(*res, productRes)
	}
	return nil
}
