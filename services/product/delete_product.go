package product

import (
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
)

var deleteProductServiceLogger = logging.GetLogger(
	"productRepo",
	"deleteProduct",
)

func (svc *productServiceImpl) DeleteProduct(req DeleteProductReq, res *DeleteProductRes) error {
	deleteRes, err := svc.repo.DeleteProduct(req.ID)
	if err != nil {
		deleteProductServiceLogger.Printf(
			"error while createProduct() caused by: %s",
			err,
		)
		return err
	}
	*res = DeleteProductRes{
		Message:      deleteRes.Message,
		RowsAffected: deleteRes.RowsAffected,
	}
	return nil
}
