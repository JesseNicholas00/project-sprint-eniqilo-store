package product

import (
	"github.com/JesseNicholas00/EniqiloStore/repos/product"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/google/uuid"
)

var createProductServiceLogger = logging.GetLogger(
	"productRepo",
	"createProduct",
)

func (svc *productServiceImpl) CreateProduct(req CreateProductReq, res *CreateProductRes) error {
	product := product.Product{
		ProductID: uuid.New().String(),
		Name:      req.Name,
		SKU:       req.SKU,
		Category:  req.Category,
		ImageUrl:  req.ImageUrl,
		Notes:     req.Notes,
		Price:     req.Price,
		Location:  req.Location,
		Available: req.Available,
	}
	savedProduct, err := svc.repo.CreateProduct(product)
	if err != nil {
		createProductServiceLogger.Printf(
			"error while createProduct() caused by: %s",
			err,
		)
		return err
	}
	*res = CreateProductRes{
		ID:        savedProduct.ProductID,
		CreatedAt: savedProduct.CreatedAt,
	}
	return nil
}
