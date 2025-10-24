package almacenseccion

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type AlmacenSeccionRepository interface {
	Create(almseccion *AlmacenSeccion) error
	GetByID(id uint) (*AlmacenSeccion, error)
	GetByAlmacen(IdAlmacen uint) ([]*AlmacenSeccion, error)
	UpdateAlmacenSeccion(id uint, updated *AlmacenSeccion) (*AlmacenSeccion, error)
	DeleteAlmacenSeccion(id uint) (string, error)
}

type almacenSeccionRepository struct {
	db *gorm.DB
}

func NewAlmacenSeccionRepository(db *gorm.DB) AlmacenSeccionRepository {
	return &almacenSeccionRepository{db}
}

func (r *almacenSeccionRepository) Create(almseccion *AlmacenSeccion) error {
	return r.db.Create(almseccion).Error
}

func (r *almacenSeccionRepository) GetByID(id uint) (*AlmacenSeccion, error) {
	var almseccion AlmacenSeccion
	err := r.db.Where("estado = ?", true).First(&almseccion, id).Error
	if err != nil {
		return nil, err
	}
	return &almseccion, nil
}

func (r *almacenSeccionRepository) GetByAlmacen(IdAlmacen uint) ([]*AlmacenSeccion, error) {
	var almsecciones []*AlmacenSeccion
	err := r.db.Where("id_almacen = ? AND estado = ?", IdAlmacen, true).Find(&almsecciones).Error
	if err != nil {
		return nil, err
	}
	return almsecciones, nil
}

func (r *almacenSeccionRepository) UpdateAlmacenSeccion(id uint, updated *AlmacenSeccion) (*AlmacenSeccion, error) {
	var almseccion AlmacenSeccion
	err := r.db.First(&almseccion, id).Error
	if err != nil {
		return nil, err
	}

	almseccion.IdAlmacen = updated.IdAlmacen
	almseccion.IdSeccion = updated.IdSeccion
	almseccion.CapacidadSeccion = updated.CapacidadSeccion
	almseccion.Estado = updated.Estado
	almseccion.FechaModificacion = updated.FechaModificacion

	err = r.db.Save(&almseccion).Error
	if err != nil {
		return nil, err
	}

	return &almseccion, nil
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
