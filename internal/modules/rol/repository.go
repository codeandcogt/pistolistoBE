package rol

import "gorm.io/gorm"

type RolRepository interface {
	Create(rol *Rol) error
	GetAll() ([]*Rol, error)
}

type rolRepository struct {
	db *gorm.DB
}

func NewRol(db *gorm.DB) RolRepository {
	return &rolRepository{db}
}

func (r *rolRepository) Create(rol *Rol) error {
	return r.db.Create(rol).Error
}

func (r *rolRepository) GetAll() ([]*Rol, error) {
	var rol []*Rol

	err := r.db.Where("estado = ?", true).Find(&rol).Error

	if err != nil {
		return nil, err
	}

	return rol, nil
}
