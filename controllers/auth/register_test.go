package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JesseNicholas00/EniqiloStore/services/auth"
	"github.com/JesseNicholas00/EniqiloStore/utils/helper"
	"github.com/JesseNicholas00/EniqiloStore/utils/unittesting"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRegisterValid(t *testing.T) {
	Convey("When given a valid register request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		userId := "dummyId"
		name := "namadepan namabelakang"
		phoneNumber := "+1-2468123123123"
		password := "password"
		accessToken := "token"

		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodPost,
			"/v1/staff/register",
			rec,
			unittesting.WithJsonPayload(map[string]interface{}{
				"phoneNumber": phoneNumber,
				"name":        name,
				"password":    password,
			}),
		)

		Convey("Should forward the request to the service layer", func() {
			expectedReq := auth.RegisterStaffReq{
				PhoneNumber: phoneNumber,
				Name:        name,
				Password:    password,
			}
			expectedRes := auth.RegisterStaffRes{
				UserId:      userId,
				PhoneNumber: phoneNumber,
				Name:        name,
				AccessToken: accessToken,
			}

			service.
				EXPECT().
				RegisterStaff(expectedReq, gomock.Any()).
				Do(func(_ auth.RegisterStaffReq, res *auth.RegisterStaffRes) {
					*res = expectedRes
				}).
				Return(nil).
				Times(1)

			unittesting.CallController(ctx, controller.registerStaff)

			Convey(
				"Should return the expected response with HTTP 201",
				func() {
					So(rec.Code, ShouldEqual, http.StatusCreated)

					expectedBody := helper.MustMarshalJson(
						map[string]interface{}{
							"message": "User registered successfully",
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

func TestRegisterInvalid(t *testing.T) {
	Convey("When given an invalid register request", t, func() {
		mockCtrl, controller, _ := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		phoneNumber := "+1-2468123123123"
		password := "password"

		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodPost,
			"/v1/staff/register",
			rec,
			unittesting.WithJsonPayload(map[string]interface{}{
				// no name
				"phoneNumber": phoneNumber,
				"password":    password,
			}),
		)

		Convey("Should return HTTP code 400", func() {
			unittesting.CallController(ctx, controller.registerStaff)
			So(rec.Code, ShouldEqual, http.StatusBadRequest)
		})
	})
}
