package service

// TODO: [XL] Actually implement auth service with callouts to hashing and DB
type AuthService interface {
	Login(username string, password string) (bool, string, error)
	Verify(token string) (bool, int, error)
}

type authService struct {}

func NewAuthService() AuthService {
	return &authService{}
}

func (h *authService) Login(username string, password string) (bool, string, error) {
	return true, "token", nil
}

func (h *authService) Verify(token string) (bool, int, error) {
	if token == "token" {
		return true, 66, nil
	}
	return false, -1, nil
}
