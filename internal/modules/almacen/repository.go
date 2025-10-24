package almacen

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type AlmacenRepository interface {
	Create(almacen *Almacen) error
	GetByID(id uint) (*Almacen, error)
	GetAllBySucursal(IdSucursal uint) ([]*Almacen, error)
	UpdateAlmacen(id uint, updated *Almacen) (*Almacen, error)
	DeleteAlmacen(id uint) (string, error)
}

type almacenRepository struct {
	db *gorm.DB
}

func NewAlmacenRepository(db *gorm.DB) AlmacenRepository {
	return &almacenRepository{db}
}

func (r *almacenRepository) Create(almacen *Almacen) error {
	return r.db.Create(almacen).Error
}

func (r *almacenRepository) GetByID(id uint) (*Almacen, error) {
	var almacen Almacen
	err := r.db.Where("estado = ?", true).First(&almacen, id).Error
	if err != nil {
		return nil, err
	}
	return &almacen, nil
}

func (r *almacenRepository) GetAllBySucursal(IdSucursal uint) ([]*Almacen, error) {
	var almacenes []*Almacen
	err := r.db.Where("id_sucursal = ? AND estado = ?", IdSucursal, true).Find(&almacenes).Error
	if err != nil {
		return nil, err
	}
	return almacenes, nil
}

func (r *almacenRepository) UpdateAlmacen(id uint, updated *Almacen) (*Almacen, error) {
	var almacen Almacen
	err := r.db.First(&almacen, id).Error
	if err != nil {
		return nil, err
	}

	almacen.Nombre = updated.Nombre
	almacen.Codigo = updated.Codigo
	almacen.TipoAlmacen = updated.TipoAlmacen
	almacen.CapacidadMaxima = updated.CapacidadMaxima
	almacen.Descripcion = updated.Descripcion
	almacen.IdSucursal = updated.IdSucursal
	almacen.Estado = updated.Estado
	almacen.FechaModificacion = updated.FechaModificacion

	err = r.db.Save(&almacen).Error
	if err != nil {
		return nil, err
	}

	return &almacen, nil
}

func (r *almacenRepository) DeleteAlmacen(id uint) (string, error) {
	result := r.db.Model(&Almacen{}).Where("id_almacen = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}
	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}
