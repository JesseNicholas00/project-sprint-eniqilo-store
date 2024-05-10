package product

import (
	"github.com/JesseNicholas00/EniqiloStore/repos/product"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
)

var updateProductLogger = logging.GetLogger("productService", "update")

func (svc *productServiceImpl) UpdateProduct(
	req UpdateProductReq,
	res *UpdateProductRes,
) error {
	product := product.Product{
		ProductID: req.Id,
		Name:      req.Name,
		SKU:       req.SKU,
		Category:  req.Category,
		ImageUrl:  req.ImageUrl,
		Notes:     req.Notes,
		Price:     req.Price,
		Location:  req.Location,
		Available: req.Available,
	}

	repoRes, err := svc.repo.UpdateProduct(product)
	if err != nil {
		updateProductLogger.Printf("could not update product: %s", err)
		return err
	}

	if repoRes.RowsAffected == 0 {
		return ErrProductNotFound
	}

	return nil
}
