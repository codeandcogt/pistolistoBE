package carrito

import (
	"gorm.io/gorm"
)

type CarritoRepository interface {
	Create(carrito *Carrito) error
	GetByID(id uint) (*Carrito, error)
	GetByClienteID(clienteId uint) (*Carrito, error)
	GetAll() ([]*Carrito, error)
	Update(id uint, carrito *Carrito) error
	Delete(id uint) error

	// CarritoItem methods
	AddItem(item *CarritoItem) error
	UpdateItem(id uint, item *CarritoItem) error
	RemoveItem(id uint) error
	GetItemsByCarritoID(carritoId uint) ([]*CarritoItem, error)
}

type carritoRepository struct {
	db *gorm.DB
}

func NewCarritoRepository(db *gorm.DB) CarritoRepository {
	return &carritoRepository{db}
}

func (r *carritoRepository) Create(carrito *Carrito) error {
	return r.db.Create(carrito).Error
}

func (r *carritoRepository) GetByID(id uint) (*Carrito, error) {
	var carrito Carrito
	err := r.db.Preload("Items", "estado = ?", true).Where("id_carrito = ? AND estado = ?", id, true).First(&carrito).Error
	if err != nil {
		return nil, err
	}
	return &carrito, nil
}

func (r *carritoRepository) GetByClienteID(clienteId uint) (*Carrito, error) {
	var carrito Carrito
	err := r.db.Preload("Items", "estado = ?", true).Where("cliente_id = ? AND estado = ? AND pedido_id IS NULL", clienteId, true).First(&carrito).Error
	if err != nil {
		return nil, err
	}
	return &carrito, nil
}

func (r *carritoRepository) GetAll() ([]*Carrito, error) {
	var carritos []*Carrito
	err := r.db.Preload("Items", "estado = ?", true).Where("estado = ?", true).Find(&carritos).Error
	if err != nil {
		return nil, err
	}
	return carritos, nil
}

func (r *carritoRepository) Update(id uint, carrito *Carrito) error {
	return r.db.Model(&Carrito{}).Where("id_carrito = ? AND estado = ?", id, true).Updates(carrito).Error
}

func (r *carritoRepository) Delete(id uint) error {
	return r.db.Model(&Carrito{}).Where("id_carrito = ? AND estado = ?", id, true).Update("estado", false).Error
}

// CarritoItem methods
func (r *carritoRepository) AddItem(item *CarritoItem) error {
	return r.db.Create(item).Error
}

func (r *carritoRepository) UpdateItem(id uint, item *CarritoItem) error {
	return r.db.Model(&CarritoItem{}).Where("id_carrito_item = ? AND estado = ?", id, true).Updates(item).Error
}

func (r *carritoRepository) RemoveItem(id uint) error {
	return r.db.Model(&CarritoItem{}).Where("id_carrito_item = ? AND estado = ?", id, true).Update("estado", false).Error
}

func (r *carritoRepository) GetItemsByCarritoID(carritoId uint) ([]*CarritoItem, error) {
	var items []*CarritoItem
	err := r.db.Where("carrito_id = ? AND estado = ?", carritoId, true).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}
