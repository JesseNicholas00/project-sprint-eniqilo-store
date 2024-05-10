package product

import (
	"github.com/JesseNicholas00/EniqiloStore/repos/product"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
)

var updateProductServiceLogger = logging.GetLogger(
	"productRepo",
	"updateProduct",
)

func (svc *productServiceImpl) UpdateProduct(req UpdateProductReq, res *UpdateProductRes) error {
	product := product.Product{
		ProductID: req.ID,
		Name:      req.Name,
		SKU:       req.SKU,
		Category:  req.Category,
		ImageUrl:  req.ImageUrl,
		Notes:     req.Notes,
		Price:     req.Price,
		Location:  req.Location,
		Available: req.Available,
	}
	updateResult, err := svc.repo.UpdateProduct(product)
	if err != nil {
		updateProductServiceLogger.Printf(
			"error while updateProduct() caused by: %s",
			err,
		)
		return err
	}
	*res = UpdateProductRes{
		Message:      updateResult.Message,
		RowsAffected: updateResult.RowsAffected,
	}
	return nil
}
