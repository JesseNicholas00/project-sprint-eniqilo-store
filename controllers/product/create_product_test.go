package product

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JesseNicholas00/EniqiloStore/services/product"
	"github.com/JesseNicholas00/EniqiloStore/utils/helper"
	"github.com/JesseNicholas00/EniqiloStore/utils/unittesting"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateProductValid(t *testing.T) {
	mockCtrl, controller, service := NewControllerWithMockedService(t)
	defer mockCtrl.Finish()

	Convey("When given a valid create product", t, func() {
		req := product.CreateProductReq{
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
		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodPost,
			"/v1/product",
			rec,
			unittesting.WithJsonPayload(helper.StructToMap(req)),
		)

		Convey("Should forward the request to the service layer", func() {
			expectedReq := req
			expectedRes := product.CreateProductRes{
				ID:        "id",
				CreatedAt: "createdAt",
			}

			service.
				EXPECT().
				CreateProduct(expectedReq, gomock.Any()).
				Do(func(_ product.CreateProductReq, res *product.CreateProductRes) {
					*res = expectedRes
				}).
				Return(nil).
				Times(1)

			unittesting.CallController(ctx, controller.createProduct)

			Convey(
				"Should return the expected response with HTTP 200",
				func() {
					So(rec.Code, ShouldEqual, http.StatusCreated)

					expectedBody := helper.MustMarshalJson(
						map[string]interface{}{
							"message": "success",
							"data":    expectedRes,
						},
					)

					So(
						rec.Body.String(),
						ShouldEqualJSON,
						string(expectedBody),
					)
				},
			)
		})
	})
}

func TestCreateProductInvalid(t *testing.T) {
	Convey("When given an invalid create product request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		req := product.CreateProductReq{
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

		Convey("On invalid request", func() {
			// missing name
			expectedReq := req
			expectedReq.Name = ""

			rec := httptest.NewRecorder()
			ctx := unittesting.CreateEchoContextFromRequest(
				http.MethodPost,
				"/v1/product",
				rec,
				unittesting.WithJsonPayload(helper.StructToMap(expectedReq)),
			)

			Convey("Should return HTTP code 400", func() {
				unittesting.CallController(ctx, controller.createProduct)
				So(rec.Code, ShouldEqual, http.StatusBadRequest)
			})
		})

		Convey("On internal server error", func() {
			expectedReq := req
			rec := httptest.NewRecorder()
			ctx := unittesting.CreateEchoContextFromRequest(
				http.MethodPost,
				"/v1/staff/register",
				rec,
				unittesting.WithJsonPayload(helper.StructToMap(expectedReq)),
			)

			service.
				EXPECT().
				CreateProduct(expectedReq, gomock.Any()).
				Return(errors.New("system error")).
				Times(1)

			Convey("Should return HTTP code 400", func() {
				unittesting.CallController(ctx, controller.createProduct)
				So(rec.Code, ShouldEqual, http.StatusInternalServerError)
			})
		})
	})
}
