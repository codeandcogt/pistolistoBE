package departamento

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type DepartamentoRepository interface {
	Create(departamento *Departamento) error
	GetByID(id uint) (*Departamento, error)
	GetAll() ([]*Departamento, error)
	DeleteDepartamento(id uint) (string, error)
}

type departamentoRepository struct {
	db *gorm.DB
}

func NewDepartamento(db *gorm.DB) DepartamentoRepository {
	return &departamentoRepository{db}
}

func (r *departamentoRepository) Create(departamento *Departamento) error {
	return r.db.Create(departamento).Error
}

func (r *departamentoRepository) GetByID(id uint) (*Departamento, error) {
	var departamento Departamento
	err := r.db.Where("estado = ?", true).First(&departamento, id).Error
	if err != nil {
		return nil, err
	}
	return &departamento, nil
}

func (r *departamentoRepository) GetAll() ([]*Departamento, error) {
	var departamento []*Departamento
	err := r.db.Where("estado = ?", true).Find(&departamento).Error

	if err != nil {
		return nil, err
	}
	return departamento, nil
}

func (r *departamentoRepository) DeleteDepartamento(id uint) (string, error) {
	result := r.db.Model(&Departamento{}).Where("id_departamento = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}
