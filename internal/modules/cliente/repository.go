package cliente

import (
	"gorm.io/gorm"
)

type ClienteRepository interface {
	Create(cliente *Cliente) error
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
