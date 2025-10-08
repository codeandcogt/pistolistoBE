package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtManager interface {
	GenerateToken(idCliente uint) (string, error)
}

type jwtManager struct {
	secret []byte
	expiry time.Duration
}

func NewJwtManager(secret []byte, expiry time.Duration) JwtManager {
	return &jwtManager{secret: secret, expiry: expiry}
}

func (j *jwtManager) GenerateToken(idCliente uint) (string, error) {
	claims := jwt.MapClaims{
		"id_cliente": idCliente,
		"exp":        time.Now().Add(time.Duration(j.expiry) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}

func (j *jwtManager) GenerateTokenAdmin(idAdministrativo uint, IdRol uint) (string, error) {
	claims := jwt.MapClaims{
		"id_admin": idAdministrativo,
		"rol":      IdRol,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}
