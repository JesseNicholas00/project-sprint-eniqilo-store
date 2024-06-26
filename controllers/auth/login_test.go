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

func TestLoginValid(t *testing.T) {
	Convey("When given a valid login request", t, func() {
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
			"/v1/staff/login",
			rec,
			unittesting.WithJsonPayload(map[string]interface{}{
				"phoneNumber": phoneNumber,
				"password":    password,
			}),
		)

		Convey("Should forward the request to the service layer", func() {
			expectedReq := auth.LoginStaffReq{
				PhoneNumber: phoneNumber,
				Password:    password,
			}
			expectedRes := auth.LoginStaffRes{
				UserId:      userId,
				PhoneNumber: phoneNumber,
				Name:        name,
				AccessToken: accessToken,
			}

			service.
				EXPECT().
				LoginStaff(expectedReq, gomock.Any()).
				Do(func(_ auth.LoginStaffReq, res *auth.LoginStaffRes) {
					*res = expectedRes
				}).
				Return(nil).
				Times(1)

			unittesting.CallController(ctx, controller.loginStaff)

			Convey(
				"Should return the expected response with HTTP 200",
				func() {
					So(rec.Code, ShouldEqual, http.StatusOK)

					expectedBody := helper.MustMarshalJson(
						map[string]interface{}{
							"message": "User logged in successfully",
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

func TestLoginInvalid(t *testing.T) {
	Convey("When given an invalid login request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		Convey("On bad request", func() {
			// wrong phone number format
			phoneNumber := "12468123123123"
			password := "password"

			rec := httptest.NewRecorder()
			ctx := unittesting.CreateEchoContextFromRequest(
				http.MethodPost,
				"/v1/staff/login",
				rec,
				unittesting.WithJsonPayload(map[string]interface{}{
					"phoneNumber": phoneNumber,
					"password":    password,
				}),
			)

			Convey("Should return HTTP code 400", func() {
				unittesting.CallController(ctx, controller.loginStaff)
				So(rec.Code, ShouldEqual, http.StatusBadRequest)
			})
		})

		phoneNumber := "+1-2468123123123"
		password := "password"
		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodPost,
			"/v1/staff/login",
			rec,
			unittesting.WithJsonPayload(map[string]interface{}{
				"phoneNumber": phoneNumber,
				"password":    password,
			}),
		)

		expectedReq := auth.LoginStaffReq{
			PhoneNumber: phoneNumber,
			Password:    password,
		}

		Convey("On user not found", func() {
			service.
				EXPECT().
				LoginStaff(expectedReq, gomock.Any()).
				Return(auth.ErrUserNotFound).
				Times(1)

			Convey(
				"Should return HTTP code 404",
				func() {
					unittesting.CallController(ctx, controller.loginStaff)
					So(rec.Code, ShouldEqual, http.StatusNotFound)
				},
			)
		})

		Convey("On invalid credentials", func() {
			service.
				EXPECT().
				LoginStaff(expectedReq, gomock.Any()).
				Return(auth.ErrInvalidCredentials).
				Times(1)

			Convey(
				"Should return HTTP code 404",
				func() {
					unittesting.CallController(ctx, controller.loginStaff)
					So(rec.Code, ShouldEqual, http.StatusBadRequest)
				},
			)
		})
	})
}
