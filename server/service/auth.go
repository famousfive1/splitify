package service

import (
	"splitify/db"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(username string, password string) (bool, error)
	Login(username string, password string) (bool, string, error)
	Verify(token string) (bool, int, error)
}

type authService struct {
	loginRepo db.LoginRepo
}

func NewAuthService(loginRepo db.LoginRepo) AuthService {
	return &authService{
		loginRepo: loginRepo,
	}
}

func (h *authService) Register(username string, password string) (bool, error) {
	passwordBytes := []byte(password)

	hashed, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}

	success, err := h.loginRepo.Insert(username, string(hashed))
	if err != nil {
		return false, err
	}

	return success, nil
}

func (h *authService) Login(username string, password string) (bool, string, error) {
	actual, err := h.loginRepo.Select(username)
	if err != nil {
		return false, "", err
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(actual.Password), []byte(password))
	if err != nil {
		return false, "", nil
	}

	token, err := getToken(username)
	return true, token, nil
}

// TODO: [M] Implement token generation and verification service for secure encoded user token
func getToken(_ string) (string, error) {
	return "token", nil
}

func (h *authService) Verify(token string) (bool, int, error) {
	if token == "token" {
		return true, 66, nil
	}
	return false, -1, nil
}
