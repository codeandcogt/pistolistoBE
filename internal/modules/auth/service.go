package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(email, password string) (*string, error)
}

type authService struct {
	repo AuthRepository
	jwt  JwtManager
}

func NewAuthService(repo AuthRepository, jwt JwtManager) AuthService {
	return &authService{repo: repo, jwt: jwt}
}

func (s *authService) Login(email, password string) (*string, error) {
	if email == "" || password == "" {
		return nil, errors.New("email y contraseña son requeridos")
	}

	client, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("usuario no encontrado")
	}

	err = bcrypt.CompareHashAndPassword([]byte(client.Contrasena), []byte(password))
	if err != nil {
		return nil, errors.New("contraseña incorrecta")
	}

	token, err := s.jwt.GenerateToken(client.IdCliente)
	if err != nil {
		return nil, errors.New("no se pudo generar token")
	}

	return &token, nil
}
