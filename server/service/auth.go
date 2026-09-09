package service

type AuthService interface {
	Login(username string, password string) (string, error)
	Verify(token string) (bool, error)
}

type authService struct {}

func NewAuthService() AuthService {
	return &authService{}
}

func (h *authService) Login(username string, password string) (string, error) {
	return "token", nil
}

func (h *authService) Verify(token string) (bool, error) {
	hash := "hash"
	return hash == "hash", nil
}
