package product

import "github.com/JesseNicholas00/EniqiloStore/utils/logging"

var createProductRepoLogger = logging.GetLogger(
	"productRepo",
	"createProduct",
)

func (repo *productRepositoryImpl) CreateProduct(
	product Product,
) (Product, error) {
	insertQuery := `INSERT INTO products(
			product_id,
			product_name,
			product_sku,
			product_category,
			product_image_url,
			product_stock,
			product_notes,
			product_price,
			product_location,
			product_is_available)
		VALUES (
			:product_id,
			:product_name,
			:product_sku,
			:product_category,
			:product_image_url,
			:product_stock,
			:product_notes,
			:product_price,
			:product_location,
			:product_is_available)
		RETURNING
			product_id,
			product_name,
			product_sku,
			product_category,
			product_image_url,
			product_stock,
			product_notes,
			product_price,
			product_location,
			product_is_available,
			created_at,
			updated_at`
	rows, err := repo.db.NamedQuery(insertQuery, product)
	res := Product{}
	if err != nil {
		createProductRepoLogger.Printf(
			"error while createProduct() caused by: %s",
			err,
		)
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&res)
		if err != nil {
			createProductRepoLogger.Printf(
				"error while createProduct() caused by: %s",
				err,
			)
			return res, err
		}
	}
	return res, nil
}
