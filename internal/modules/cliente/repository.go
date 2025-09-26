package cliente

import (
	"gorm.io/gorm"
)

type ClienteRepository interface {
	Create(cliente *Cliente) error
	GetByID(id uint) (*Cliente, error)
	GetAll() ([]*Cliente, error)
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
	err := r.db.First(&cliente, id).Error
	if err != nil {
		return nil, err
	}
	return &cliente, nil
}

func (r *clienteRepository) GetAll() ([]*Cliente, error) {
	var clientes []*Cliente
	err := r.db.Find(&clientes).Error

	if err != nil {
		return nil, err
	}
	return clientes, nil
}
