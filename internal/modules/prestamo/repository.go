package prestamo

import (
	"gorm.io/gorm"
)

type PrestamoRepository interface {
	Create(prestamo *Prestamo) error
	GetByID(id uint) (*Prestamo, error)
	GetAll() ([]*Prestamo, error)
	GetByClienteID(clienteId uint) ([]*Prestamo, error)
	Update(id uint, prestamo *Prestamo) error
	Delete(id uint) error
}

type prestamoRepository struct {
	db *gorm.DB
}

func NewPrestamoRepository(db *gorm.DB) PrestamoRepository {
	return &prestamoRepository{db}
}

func (r *prestamoRepository) Create(prestamo *Prestamo) error {
	return r.db.Create(prestamo).Error
}

func (r *prestamoRepository) GetByID(id uint) (*Prestamo, error) {
	var prestamo Prestamo
	err := r.db.Where("id_prestamo = ? AND estado = ?", id, true).First(&prestamo).Error
	if err != nil {
		return nil, err
	}
	return &prestamo, nil
}

func (r *prestamoRepository) GetAll() ([]*Prestamo, error) {
	var prestamos []*Prestamo
	err := r.db.Where("estado = ?", true).Order("fecha_creacion DESC").Find(&prestamos).Error
	if err != nil {
		return nil, err
	}
	return prestamos, nil
}

func (r *prestamoRepository) GetByClienteID(clienteId uint) ([]*Prestamo, error) {
	var prestamos []*Prestamo
	err := r.db.Where("cliente_id = ? AND estado = ?", clienteId, true).Order("fecha_prestamo DESC").Find(&prestamos).Error
	if err != nil {
		return nil, err
	}
	return prestamos, nil
}

func (r *prestamoRepository) Update(id uint, prestamo *Prestamo) error {
	return r.db.Model(&Prestamo{}).Where("id_prestamo = ? AND estado = ?", id, true).Updates(prestamo).Error
}

func (r *prestamoRepository) Delete(id uint) error {
	return r.db.Model(&Prestamo{}).Where("id_prestamo = ? AND estado = ?", id, true).Update("estado", false).Error
}
