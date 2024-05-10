package product

import (
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
)

var deleteProductServiceLogger = logging.GetLogger(
	"productRepo",
	"deleteProduct",
)

func (svc *productServiceImpl) DeleteProduct(req DeleteResultReq, res *DeleteResultRes) error {
	deleteRes, err := svc.repo.DeleteProduct(req.ID)
	if err != nil {
		deleteProductServiceLogger.Printf(
			"error while deleteProduct() caused by: %s",
			err,
		)
		return err
	}
	*res = DeleteResultRes{
		Message:      deleteRes.Message,
		RowsAffected: deleteRes.RowsAffected,
	}
	return nil
}
