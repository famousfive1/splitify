package service

import "fmt"

type AuthService interface {
	Login(username string, password string) (string, error)
	Verify(token string) (int, error)
}

type authService struct {}

func NewAuthService() AuthService {
	return &authService{}
}

func (h *authService) Login(username string, password string) (string, error) {
	return "token", nil
}

func (h *authService) Verify(token string) (int, error) {
	if token == "token" {
		return 66, nil
	}
	return -1, fmt.Errorf("Invliad session ID")
}
