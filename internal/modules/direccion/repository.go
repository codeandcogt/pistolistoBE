package direccion

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type DireccionRepository interface {
	Create(direccion *Direccion) error
	GetByID(id uint) (*Direccion, error)
	GetAll() ([]*Direccion, error)
	Update(direccion *Direccion) (string, error)
	Delete(id uint) (string, error)
}

type direccionRepository struct {
	db *gorm.DB
}

func NewDireccionRepository(db *gorm.DB) DireccionRepository {
	return &direccionRepository{db}
}

func (r *direccionRepository) Create(direccion *Direccion) error {
	return r.db.Create(direccion).Error
}

func (r *direccionRepository) GetByID(id uint) (*Direccion, error) {
	var direccion Direccion
	err := r.db.Where("estado = ?", true).First(&direccion, id).Error
	if err != nil {
		return nil, err
	}
	return &direccion, nil
}

func (r *direccionRepository) GetAll() ([]*Direccion, error) {
	var direcciones []*Direccion
	err := r.db.Where("estado = ?", true).Find(&direcciones).Error
	if err != nil {
		return nil, err
	}
	return direcciones, nil
}

func (r *direccionRepository) Update(direccion *Direccion) (string, error) {
	result := r.db.Model(&Direccion{}).
		Where("id_direccion = ? AND estado = ?", direccion.IdDireccion, true).
		Updates(direccion)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_UPDATED, nil
}

func (r *direccionRepository) Delete(id uint) (string, error) {
	result := r.db.Model(&Direccion{}).Where("id_direccion = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}
