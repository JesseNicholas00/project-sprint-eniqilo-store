package product

import (
	"testing"

	"github.com/JesseNicholas00/EniqiloStore/repos/product"
	unittesting "github.com/JesseNicholas00/EniqiloStore/utils/unittesting"
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateProduct(t *testing.T) {
	Convey("When create product", t, func() {
		mockCtrl, service, mockedRepo := NewWithMockedRepo(t)
		defer mockCtrl.Finish()

		req := CreateProductReq{
			Name:      "name",
			SKU:       "sku",
			Category:  "Clothing",
			ImageUrl:  "http://www.google.com",
			Notes:     "notes",
			Price:     1,
			Stock:     1,
			Location:  "location",
			Available: true,
		}

		unittesting.FixNextUuid()
		uuid := uuid.New().String()
		repoReq := product.Product{
			ProductID: uuid,
			Name:      "name",
			SKU:       "sku",
			Category:  "Clothing",
			ImageUrl:  "http://www.google.com",
			Notes:     "notes",
			Price:     1,
			Stock:     1,
			Location:  "location",
			Available: true,
		}
		repoRes := product.Product{
			ProductID: uuid,
			Name:      "name",
			SKU:       "sku",
			Category:  "Clothing",
			ImageUrl:  "http://www.google.com",
			Notes:     "notes",
			Price:     1,
			Stock:     1,
			Location:  "location",
			Available: true,
			CreatedAt: "now",
			UpdatedAt: "now",
		}

		unittesting.FixNextUuid()
		Convey("If the request is valid", func() {
			mockedRepo.EXPECT().
				CreateProduct(repoReq).
				Return(repoRes, nil).
				Times(1)

			res := CreateProductRes{}
			err := service.CreateProduct(req, &res)
			Convey(
				"Should return nil and write the correct result to res",
				func() {
					So(err, ShouldBeNil)
					So(res.ID, ShouldEqual, uuid)
					So(res.CreatedAt, ShouldEqual, repoRes.CreatedAt)
				},
			)
		})
	})
}
