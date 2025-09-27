package cliente

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type ClienteRepository interface {
	Create(cliente *Cliente) error
	GetByID(id uint) (*Cliente, error)
	GetAll() ([]*Cliente, error)
	DeleteCliente(id uint) (string, error)
}

type clienteRepository struct {
	db *gorm.DB
}

func NewClient(db *gorm.DB) ClienteRepository {
	return &clienteRepository{db}
}

func (r *clienteRepository) Create(client *Cliente) error {
	return r.db.Create(client).Error
}

func (r *clienteRepository) GetByID(id uint) (*Cliente, error) {
	var cliente Cliente
	err := r.db.Where("estado = ?", true).First(&cliente, id).Error
	if err != nil {
		return nil, err
	}
	return &cliente, nil
}

func (r *clienteRepository) GetAll() ([]*Cliente, error) {
	var clientes []*Cliente
	err := r.db.Where("estado = ?", true).Find(&clientes).Error

	if err != nil {
		return nil, err
	}
	return clientes, nil
}

func (r *clienteRepository) DeleteCliente(id uint) (string, error) {
	result := r.db.Model(&Cliente{}).Where("id_cliente = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}
