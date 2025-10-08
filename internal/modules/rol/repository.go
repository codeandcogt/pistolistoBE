package rol

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type RolRepository interface {
	Create(rol *Rol) error
	GetAll() ([]*Rol, error)
	GetByID(id uint) (*Rol, error)
	Delete(id uint) (string, error)
	Update(rol *Rol) (string, error)
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

func (r *rolRepository) GetByID(id uint) (*Rol, error) {
	var rol Rol
	err := r.db.Where("estado = ?", true).First(&rol, id).Error
	if err != nil {
		return nil, err
	}
	return &rol, nil
}

func (r *rolRepository) Delete(id uint) (string, error) {
	result := r.db.Model(&Rol{}).Where("id_rol = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}

func (r *rolRepository) Update(rol *Rol) (string, error) {
	result := r.db.Model(&Rol{}).
		Where("id_rol = ? AND estado = ?", rol.IdRol, true).
		Updates(rol)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_UPDATED, nil
}
