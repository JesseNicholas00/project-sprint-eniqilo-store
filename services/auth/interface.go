package auth

//go:generate mockgen -destination mock_test.go -package auth_test . AuthService

type AuthService interface {
	RegisterStaff(req RegisterStaffReq, res *RegisterStaffRes) error
	LoginStaff(req LoginStaffReq, res *LoginStaffRes) error
}
