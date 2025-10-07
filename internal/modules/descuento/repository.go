package descuento

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type DescuentoRepository interface {
	Create(descuento *Descuento) error
	GetByID(id uint) (*Descuento, error)
	GetAll() ([]*Descuento, error)
	DeleteDescuento(id uint) (string, error)
}

type descuentoRepository struct {
	db *gorm.DB
}

func NewDescuento(db *gorm.DB) DescuentoRepository {
	return &descuentoRepository{db}
}

func (r *descuentoRepository) Create(descuento *Descuento) error {
	return r.db.Create(descuento).Error
}

func (r *descuentoRepository) GetByID(id uint) (*Descuento, error) {
	var descuento Descuento
	err := r.db.Where("estado = ?", true).First(&descuento, id).Error
	if err != nil {
		return nil, err
	}
	return &descuento, nil
}

func (r *descuentoRepository) GetAll() ([]*Descuento, error) {
	var descuento []*Descuento
	err := r.db.Where("estado = ?", true).Find(&descuento).Error

	if err != nil {
		return nil, err
	}
	return descuento, nil
}

func (r *descuentoRepository) DeleteDescuento(id uint) (string, error) {
	result := r.db.Model(&Descuento{}).Where("id_descuento = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}
