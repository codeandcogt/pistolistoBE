package producto

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type ProductoRepository interface {
	Create(producto *Producto) error
	GetByID(id uint) (*Producto, error)
	GetAll() ([]*Producto, error)
	Update(producto *Producto) (string, error)
	Delete(id uint) (string, error)
}

type productoRepository struct {
	db *gorm.DB
}

func NewProductoRepository(db *gorm.DB) ProductoRepository {
	return &productoRepository{db}
}

func (r *productoRepository) Create(producto *Producto) error {
	return r.db.Create(producto).Error
}

func (r *productoRepository) GetByID(id uint) (*Producto, error) {
	var producto Producto
	err := r.db.Where("estado = ?", true).First(&producto, id).Error
	if err != nil {
		return nil, err
	}
	return &producto, nil
}

func (r *productoRepository) GetAll() ([]*Producto, error) {
	var productos []*Producto
	err := r.db.
		Preload("Articulo", "estado = ?", true).
		Where("estado = ?", true).
		Find(&productos).Error
	if err != nil {
		return nil, err
	}
	return productos, nil
}

func (r *productoRepository) Update(producto *Producto) (string, error) {
	result := r.db.Model(&Producto{}).
		Where("id_producto = ? AND estado = ?", producto.IdProducto, true).
		Updates(producto)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_UPDATED, nil
}

func (r *productoRepository) Delete(id uint) (string, error) {
	result := r.db.Model(&Producto{}).Where("id_producto = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}
