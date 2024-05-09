package auth

type authServiceImpl struct {
}

func NewAuthService() AuthService {
	return &authServiceImpl{}
}
