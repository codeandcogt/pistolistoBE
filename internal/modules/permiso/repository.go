package permiso

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type PermisoRepository interface {
	Create(permiso *Permiso) error
	GetAll() ([]*Permiso, error)
	GetByID(id uint) (*Permiso, error)
	Delete(id uint) (string, error)
	Update(permiso *Permiso) (string, error)
}

type permisoRepository struct {
	db *gorm.DB
}

func NewPermisoRepository(db *gorm.DB) PermisoRepository {
	return &permisoRepository{db}
}

func (r *permisoRepository) Create(rol *Permiso) error {
	return r.db.Create(rol).Error
}

func (r *permisoRepository) GetAll() ([]*Permiso, error) {
	var permiso []*Permiso

	err := r.db.Where("estado = ?", true).Find(&permiso).Error

	if err != nil {
		return nil, err
	}

	return permiso, nil
}

func (r *permisoRepository) GetByID(id uint) (*Permiso, error) {
	var permiso Permiso
	err := r.db.Where("estado = ?", true).First(&permiso, id).Error
	if err != nil {
		return nil, err
	}
	return &permiso, nil
}

func (r *permisoRepository) Delete(id uint) (string, error) {
	result := r.db.Model(&Permiso{}).Where("id_permiso = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}

func (r *permisoRepository) Update(permiso *Permiso) (string, error) {
	result := r.db.Model(&Permiso{}).
		Where("id_permiso = ? AND estado = ?", permiso.IdPermiso, true).
		Updates(permiso)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_UPDATED, nil
}
