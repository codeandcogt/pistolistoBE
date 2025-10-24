package seccion

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type SeccionRepository interface {
	Create(seccion *Seccion) error
	GetByID(id uint) (*Seccion, error)
	GetAll() ([]*Seccion, error)
	UpdateSeccion(id uint, updated *Seccion) (*Seccion, error)
	DeleteSeccion(id uint) (string, error)
}

type seccionRepository struct {
	db *gorm.DB
}

func NewSeccionRepository(db *gorm.DB) SeccionRepository {
	return &seccionRepository{db}
}

func (r *seccionRepository) Create(seccion *Seccion) error {
	return r.db.Create(seccion).Error
}

func (r *seccionRepository) GetByID(id uint) (*Seccion, error) {
	var seccion Seccion
	err := r.db.Where("estado = ?", true).First(&seccion, id).Error
	if err != nil {
		return nil, err
	}
	return &seccion, nil
}

func (r *seccionRepository) GetAll() ([]*Seccion, error) {
	var secciones []*Seccion
	err := r.db.Where("estado = ?", true).Find(&secciones).Error
	if err != nil {
		return nil, err
	}
	return secciones, nil
}

func (r *seccionRepository) UpdateSeccion(id uint, updated *Seccion) (*Seccion, error) {
	var seccion Seccion
	err := r.db.First(&seccion, id).Error
	if err != nil {
		return nil, err
	}

	seccion.Nombre = updated.Nombre
	seccion.Descripcion = updated.Descripcion
	seccion.Codigo = updated.Codigo
	seccion.TipoSeccion = updated.TipoSeccion
	seccion.Estado = updated.Estado
	seccion.FechaModificacion = updated.FechaModificacion

	err = r.db.Save(&seccion).Error
	if err != nil {
		return nil, err
	}

	return &seccion, nil
}

func (r *seccionRepository) DeleteSeccion(id uint) (string, error) {
	result := r.db.Model(&Seccion{}).Where("id_seccion = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}
	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}
