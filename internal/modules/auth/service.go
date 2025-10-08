package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(email, password string) (*string, error)
	LoginAdmin(email, password string) (*string, error)
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

	estado := true
	logEntry := &LogLoginCliente{
		IdCliente: int(client.IdCliente),
		Estado:    &estado,
		Exito:     false,
	}

	err = bcrypt.CompareHashAndPassword([]byte(client.Contrasena), []byte(password))
	if err != nil {
		s.repo.LogLoginCliente(logEntry)
		return nil, errors.New("contraseña incorrecta")
	}

	token, err := s.jwt.GenerateToken(client.IdCliente)
	if err != nil {
		s.repo.LogLoginCliente(logEntry)
		return nil, errors.New("no se pudo generar token")
	}

	logEntry.Exito = true
	s.repo.LogLoginCliente(logEntry)
	return &token, nil
}

func (s *authService) LoginAdmin(email, password string) (*string, error) {
	// if email == "" || password == "" {
	// 	return nil, errors.New("email y contraseña son requeridos")
	// }

	// admin, err := s.repo.FindByEmailAdmin(email)
	// if err != nil {
	// 	return nil, errors.New("usuario no encontrado")
	// }

	// estado := true
	// logEntry := &LogLoginCliente{
	// 	IdCliente: int(admin.IdAdministrativo),
	// 	Estado:    &estado,
	// 	Exito:     false,
	// }

	// err = bcrypt.CompareHashAndPassword([]byte(client.Contrasena), []byte(password))
	// if err != nil {
	// 	s.repo.LogLoginCliente(logEntry)
	// 	return nil, errors.New("contraseña incorrecta")
	// }

	// token, err := s.jwt.GenerateToken(client.IdCliente)
	// if err != nil {
	// 	s.repo.LogLoginCliente(logEntry)
	// 	return nil, errors.New("no se pudo generar token")
	// }

	// logEntry.Exito = true
	// s.repo.LogLoginCliente(logEntry)
	return nil, nil
}
