package rolpermiso

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type RolPermisoRepository interface {
	Create(rolPermiso *RolPermiso) error
	GetAll() ([]*RolPermiso, error)
	GetByID(id uint) (*RolPermiso, error)
	Update(rolPermiso *RolPermiso) (string, error)
	Delete(id uint) (string, error)
}

type rolPermisoRepository struct {
	db *gorm.DB
}

func NewRolPermiso(db *gorm.DB) RolPermisoRepository {
	return &rolPermisoRepository{db}
}

func (r *rolPermisoRepository) Create(rolPermiso *RolPermiso) error {
	return r.db.Create(rolPermiso).Error
}

func (r *rolPermisoRepository) GetAll() ([]*RolPermiso, error) {
	var rolPermiso []*RolPermiso

	err := r.db.Where("estado = ?", true).Find(&rolPermiso).Error

	if err != nil {
		return nil, err
	}

	return rolPermiso, nil
}

func (r *rolPermisoRepository) GetByID(id uint) (*RolPermiso, error) {
	var rolPermiso RolPermiso
	err := r.db.Where("estado = ?", true).First(&rolPermiso, id).Error
	if err != nil {
		return nil, err
	}
	return &rolPermiso, nil
}

func (r *rolPermisoRepository) Update(rolPermiso *RolPermiso) (string, error) {
	result := r.db.Model(&RolPermiso{}).
		Where("id_rol_permiso = ? AND estado = ?", rolPermiso.IdRolPermiso, true).
		Updates(rolPermiso)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_UPDATED, nil
}

func (r *rolPermisoRepository) Delete(id uint) (string, error) {
	result := r.db.Model(&RolPermiso{}).Where("id_rol_permiso = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}
