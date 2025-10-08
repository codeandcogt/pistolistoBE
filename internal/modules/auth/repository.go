package auth

import (
	"pistolistoBE/internal/modules/administrativo"
	"pistolistoBE/internal/modules/cliente"

	"gorm.io/gorm"
)

type AuthRepository interface {
	FindByEmail(email string) (*cliente.Cliente, error)
	LogLoginCliente(log *LogLoginCliente) error
	FindByEmailAdmin(email string) (*administrativo.Administrativo, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) FindByEmail(email string) (*cliente.Cliente, error) {
	var client cliente.Cliente
	if err := r.db.Where("email = ?", email).First(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *authRepository) LogLoginCliente(log *LogLoginCliente) error {
	return r.db.Create(log).Error
}

func (r *authRepository) FindByEmailAdmin(email string) (*administrativo.Administrativo, error) {
	var admin administrativo.Administrativo
	if err := r.db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
