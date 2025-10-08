package administrativo

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type AdminRepository interface {
	Create(admin *Administrativo) error
	GetByID(id uint) (*Administrativo, error)
	GetAll() ([]*Administrativo, error)
	Delete(id uint) (string, error)
	FirstAuth(id uint) (string, error)
	Update(admin *Administrativo) (string, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) Create(rol *Administrativo) error {
	return r.db.Create(rol).Error
}

func (r *adminRepository) GetAll() ([]*Administrativo, error) {
	var permiso []*Administrativo

	err := r.db.Where("estado = ?", true).Find(&permiso).Error

	if err != nil {
		return nil, err
	}

	return permiso, nil
}

func (r *adminRepository) GetByID(id uint) (*Administrativo, error) {
	var admin Administrativo
	err := r.db.Where("estado = ?", true).First(&admin, id).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) Delete(id uint) (string, error) {
	result := r.db.Model(&Administrativo{}).Where("id_permiso = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}

func (r *adminRepository) Update(permiso *Administrativo) (string, error) {
	result := r.db.Model(&Administrativo{}).
		Where("id_permiso = ? AND estado = ?", permiso.IdAdministrativo, true).
		Updates(permiso)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_UPDATED, nil
}

func (r *adminRepository) FirstAuth(id uint) (string, error) {
	result := r.db.Model(&Administrativo{}).
		Where("id_administrativo = ? AND primer_login = ? AND estado = ?", id, false, true).
		Update("primer_login", true)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	// Verificar si se actualizó algún registro
	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, result.Error
	}

	return common.SUCCESS_UPDATED, nil
}
