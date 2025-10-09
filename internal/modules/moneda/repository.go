package moneda

import (
	"gorm.io/gorm"
)

type MonedaRepository interface {
	Create(moneda *Moneda) error
	GetByID(id uint) (*Moneda, error)
	GetAll() ([]*Moneda, error)
	Update(id uint, moneda *Moneda) error
	Delete(id uint) error
	GetByCodigo(codigo string) (*Moneda, error)
}

type monedaRepository struct {
	db *gorm.DB
}

func NewMonedaRepository(db *gorm.DB) MonedaRepository {
	return &monedaRepository{db}
}

func (r *monedaRepository) Create(moneda *Moneda) error {
	return r.db.Create(moneda).Error
}

func (r *monedaRepository) GetByID(id uint) (*Moneda, error) {
	var moneda Moneda
	err := r.db.Where("id_moneda = ? AND estado = ?", id, true).First(&moneda).Error
	if err != nil {
		return nil, err
	}
	return &moneda, nil
}

func (r *monedaRepository) GetAll() ([]*Moneda, error) {
	var monedas []*Moneda
	err := r.db.Where("estado = ?", true).Find(&monedas).Error
	if err != nil {
		return nil, err
	}
	return monedas, nil
}

func (r *monedaRepository) Update(id uint, moneda *Moneda) error {
	return r.db.Model(&Moneda{}).Where("id_moneda = ? AND estado = ?", id, true).Updates(moneda).Error
}

func (r *monedaRepository) Delete(id uint) error {
	return r.db.Model(&Moneda{}).Where("id_moneda = ? AND estado = ?", id, true).Update("estado", false).Error
}

func (r *monedaRepository) GetByCodigo(codigo string) (*Moneda, error) {
	var moneda Moneda
	err := r.db.Where("codigo = ? AND estado = ?", codigo, true).First(&moneda).Error
	if err != nil {
		return nil, err
	}
	return &moneda, nil
}
