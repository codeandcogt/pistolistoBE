package almacenseccion

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type AlmacenSeccionRepository interface {
	Create(seccion *AlmacenSeccion) error
	GetByID(id uint) (*AlmacenSeccion, error)
	GetByAlmacen(idAlmacen uint) ([]*AlmacenSeccion, error)
	UpdateAlmacenSeccion(id uint, updated *AlmacenSeccion) (*AlmacenSeccion, error)
	DeleteAlmacenSeccion(id uint) (string, error)
}

type almacenSeccionRepository struct {
	db *gorm.DB
}

func NewAlmacenSeccionRepository(db *gorm.DB) AlmacenSeccionRepository {
	return &almacenSeccionRepository{db}
}

func (r *almacenSeccionRepository) Create(seccion *AlmacenSeccion) error {
	return r.db.Create(seccion).Error
}

func (r *almacenSeccionRepository) GetByID(id uint) (*AlmacenSeccion, error) {
	var seccion AlmacenSeccion
	err := r.db.Where("estado = ?", true).First(&seccion, id).Error
	if err != nil {
		return nil, err
	}
	return &seccion, nil
}

func (r *almacenSeccionRepository) GetByAlmacen(idAlmacen uint) ([]*AlmacenSeccion, error) {
	var secciones []*AlmacenSeccion
	err := r.db.Where("id_almacen = ? AND estado = ?", idAlmacen, true).Find(&secciones).Error
	if err != nil {
		return nil, err
	}
	return secciones, nil
}

func (r *almacenSeccionRepository) UpdateAlmacenSeccion(id uint, updated *AlmacenSeccion) (*AlmacenSeccion, error) {
	var seccion AlmacenSeccion
	err := r.db.First(&seccion, id).Error
	if err != nil {
		return nil, err
	}

	seccion.IdAlmacen = updated.IdAlmacen
	seccion.IdSeccion = updated.IdSeccion
	seccion.CapacidadSeccion = updated.CapacidadSeccion
	seccion.Estado = updated.Estado
	seccion.FechaModificacion = updated.FechaModificacion

	err = r.db.Save(&seccion).Error
	if err != nil {
		return nil, err
	}

	return &seccion, nil
}

func (r *almacenSeccionRepository) DeleteAlmacenSeccion(id uint) (string, error) {
	result := r.db.Model(&AlmacenSeccion{}).Where("id_almacen_seccion = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}
	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}
