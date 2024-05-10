package auth

type AuthService interface {
	RegisterStaff(req RegisterStaffReq, res *RegisterStaffRes) error
	LoginStaff(req LoginStaffReq, res *LoginStaffRes) error
	GetSessionFromToken(
		req GetSessionFromTokenReq,
		res *GetSessionFromTokenRes,
	) error
}
