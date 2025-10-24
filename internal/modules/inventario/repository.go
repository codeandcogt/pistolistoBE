package inventario

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type InventarioRepository interface {
	Create(inventario *Inventario) error
	GetByID(id uint) (*Inventario, error)
	GetAllByAlmacen(IdAlmacen uint) ([]*Inventario, error)
	GetByTipo(IdAlmacen uint, tipoItem string) ([]*Inventario, error)
	GetByEstado(IdAlmacen uint, estadoInventario string) ([]*Inventario, error)
	GetProductosDisponibles(IdAlmacen uint) ([]*Inventario, error)
	GetArticulosEmpenados(IdAlmacen uint) ([]*Inventario, error)
	UpdateInventario(id uint, updated *Inventario) (*Inventario, error)
	DeleteInventario(id uint) (string, error)
}

type inventarioRepository struct {
	db *gorm.DB
}

func NewInventarioRepository(db *gorm.DB) InventarioRepository {
	return &inventarioRepository{db}
}

func (r *inventarioRepository) Create(inventario *Inventario) error {
	return r.db.Create(inventario).Error
}

func (r *inventarioRepository) GetByID(id uint) (*Inventario, error) {
	var inventario Inventario
	err := r.db.Where("estado = ?", true).
		Preload("Almacen").
		Preload("Producto").
		Preload("Articulo").
		First(&inventario, id).Error
	if err != nil {
		return nil, err
	}
	return &inventario, nil
}

func (r *inventarioRepository) GetAllByAlmacen(IdAlmacen uint) ([]*Inventario, error) {
	var inventarios []*Inventario
	err := r.db.Where("id_almacen = ? AND estado = ?", IdAlmacen, true).
		Preload("Almacen").
		Preload("Producto").
		Preload("Articulo").
		Find(&inventarios).Error
	if err != nil {
		return nil, err
	}
	return inventarios, nil
}

func (r *inventarioRepository) GetByTipo(IdAlmacen uint, tipoItem string) ([]*Inventario, error) {
	var inventarios []*Inventario
	err := r.db.Where("id_almacen = ? AND tipo_item = ? AND estado = ?", IdAlmacen, tipoItem, true).
		Preload("Almacen").
		Preload("Producto").
		Preload("Articulo").
		Find(&inventarios).Error
	if err != nil {
		return nil, err
	}
	return inventarios, nil
}

func (r *inventarioRepository) GetByEstado(IdAlmacen uint, estadoInventario string) ([]*Inventario, error) {
	var inventarios []*Inventario
	err := r.db.Where("id_almacen = ? AND estado_inventario = ? AND estado = ?", IdAlmacen, estadoInventario, true).
		Preload("Almacen").
		Preload("Producto").
		Preload("Articulo").
		Find(&inventarios).Error
	if err != nil {
		return nil, err
	}
	return inventarios, nil
}

func (r *inventarioRepository) GetProductosDisponibles(IdAlmacen uint) ([]*Inventario, error) {
	var inventarios []*Inventario
	err := r.db.Where("id_almacen = ? AND tipo_item = ? AND estado_inventario = ? AND estado = ?",
		IdAlmacen, "producto", "disponible", true).
		Preload("Almacen").
		Preload("Producto").
		Find(&inventarios).Error
	if err != nil {
		return nil, err
	}
	return inventarios, nil
}

func (r *inventarioRepository) GetArticulosEmpenados(IdAlmacen uint) ([]*Inventario, error) {
	var inventarios []*Inventario
	err := r.db.Where("id_almacen = ? AND tipo_item = ? AND estado_inventario = ? AND estado = ?",
		IdAlmacen, "articulo", "empeñado", true).
		Preload("Almacen").
		Preload("Articulo").
		Find(&inventarios).Error
	if err != nil {
		return nil, err
	}
	return inventarios, nil
}

func (r *inventarioRepository) UpdateInventario(id uint, updated *Inventario) (*Inventario, error) {
	var inventario Inventario
	err := r.db.First(&inventario, id).Error
	if err != nil {
		return nil, err
	}

	inventario.Codigo = updated.Codigo
	inventario.TipoItem = updated.TipoItem
	inventario.IdProducto = updated.IdProducto
	inventario.IdArticulo = updated.IdArticulo
	inventario.IdAlmacen = updated.IdAlmacen
	inventario.Cantidad = updated.Cantidad
	inventario.Ubicacion = updated.Ubicacion
	inventario.EstadoInventario = updated.EstadoInventario
	inventario.FechaIngreso = updated.FechaIngreso
	inventario.FechaSalida = updated.FechaSalida
	inventario.Observaciones = updated.Observaciones
	inventario.Estado = updated.Estado
	inventario.FechaModificacion = updated.FechaModificacion

	err = r.db.Save(&inventario).Error
	if err != nil {
		return nil, err
	}

	return &inventario, nil
}

func (r *inventarioRepository) DeleteInventario(id uint) (string, error) {
	result := r.db.Model(&Inventario{}).Where("id_inventario = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}
	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}
