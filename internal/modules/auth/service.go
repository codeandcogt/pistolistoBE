package auth

import (
	"errors"
	"pistolistoBE/internal/common"
	"pistolistoBE/internal/modules/administrativo"
	"pistolistoBE/internal/modules/cliente"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(email, password string) (*DataCliente, error)
	LoginAdmin(email, password string) (*DataAdmin, error)
}

type authService struct {
	repo AuthRepository
	jwt  JwtManager
}

type DataAdmin struct {
	Admin *administrativo.Administrativo `json:"admin"`
	Token string                         `json:"token"`
}

type DataCliente struct {
	Cliente *cliente.Cliente `json:"cliente"`
	Token   string           `json:"token"`
}

func NewAuthService(repo AuthRepository, jwt JwtManager) AuthService {
	return &authService{repo: repo, jwt: jwt}
}

func (s *authService) Login(email, password string) (*DataCliente, error) {
	if email == "" || password == "" {
		return nil, errors.New(common.ERR_REQUIRED_FIELD)
	}

	client, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New(common.ERR_NOT_FOUND)
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
		return nil, errors.New(common.ERR_INVALID_LOGIN)
	}

	token, err := s.jwt.GenerateToken(client.IdCliente)
	if err != nil {
		s.repo.LogLoginCliente(logEntry)
		return nil, errors.New(common.ERR_INVALID_LOGIN)
	}

	logEntry.Exito = true
	s.repo.LogLoginCliente(logEntry)

	dataCliente := &DataCliente{
		Cliente: client,
		Token:   token,
	}
	return dataCliente, nil
}

func (s *authService) LoginAdmin(email, password string) (*DataAdmin, error) {
	if email == "" || password == "" {
		return nil, errors.New(common.ERR_REQUIRED_FIELD)
	}

	admin, err := s.repo.FindByEmailAdmin(email)

	if err != nil {
		return nil, errors.New(common.ERR_NOT_FOUND)
	}

	logEntry := &LogLoginAdmin{
		IdAdministrativo: admin.IdAdministrativo,
		Exito:            false,
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Contrasenia), []byte(password))
	if err != nil {
		s.repo.LogLoginAdmin(logEntry)
		return nil, errors.New(common.ERR_INVALID_LOGIN)
	}

	token, err := s.jwt.GenerateTokenAdmin(admin.IdAdministrativo, admin.IdRol)
	if err != nil {
		s.repo.LogLoginAdmin(logEntry)
		return nil, errors.New(common.ERR_INVALID_LOGIN)
	}

	logEntry.Exito = true
	s.repo.LogLoginAdmin(logEntry)
	data := &DataAdmin{
		Admin: admin,
		Token: token,
	}
	return data, nil
}
