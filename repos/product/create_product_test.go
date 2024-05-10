//go:build integration
// +build integration

package product_test

import (
	"testing"

	"github.com/JesseNicholas00/EniqiloStore/repos/product"
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateProduct(t *testing.T) {
	Convey("When creating new product", t, func() {
		repo := NewWithTestDatabase(t)

		reqProduct := product.Product{
			ProductID: uuid.New().String(),
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

		savedProduct, err := repo.CreateProduct(reqProduct)
		Convey("Should return the created staff with the same data", func() {
			So(err, ShouldBeNil)
			So(savedProduct.ProductID, ShouldEqual, reqProduct.ProductID)
			So(savedProduct.Name, ShouldEqual, reqProduct.Name)
			So(savedProduct.SKU, ShouldEqual, reqProduct.SKU)
			So(savedProduct.Category, ShouldEqual, reqProduct.Category)
			So(savedProduct.ImageUrl, ShouldEqual, reqProduct.ImageUrl)
			So(savedProduct.Notes, ShouldEqual, reqProduct.Notes)
			So(savedProduct.Price, ShouldEqual, reqProduct.Price)
			So(savedProduct.Stock, ShouldEqual, reqProduct.Stock)
			So(savedProduct.Location, ShouldEqual, reqProduct.Location)
			So(savedProduct.Available, ShouldEqual, reqProduct.Available)
			So(savedProduct.CreatedAt, ShouldNotBeNil)
			So(savedProduct.UpdatedAt, ShouldNotBeNil)
		})
	})
}
