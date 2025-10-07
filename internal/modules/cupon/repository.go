package cupon

import (
	"gorm.io/gorm"
)

type CuponRepository interface {
	Create(cupon *Cupon) error
	GetByID(id uint) (*Cupon, error)
	GetAll() ([]*Cupon, error)
	Update(id uint, cupon *Cupon) error
	Delete(id uint) error
	GetByCodigo(codigo string) (*Cupon, error)
}

type cuponRepository struct {
	db *gorm.DB
}

func NewCuponRepository(db *gorm.DB) CuponRepository {
	return &cuponRepository{db}
}

func (r *cuponRepository) Create(cupon *Cupon) error {
	return r.db.Create(cupon).Error
}

func (r *cuponRepository) GetByID(id uint) (*Cupon, error) {
	var cupon Cupon
	err := r.db.Where("id_cupon = ? AND estado = ?", id, true).First(&cupon).Error
	if err != nil {
		return nil, err
	}
	return &cupon, nil
}

func (r *cuponRepository) GetAll() ([]*Cupon, error) {
	var cupones []*Cupon
	err := r.db.Where("estado = ?", true).Find(&cupones).Error
	if err != nil {
		return nil, err
	}
	return cupones, nil
}

func (r *cuponRepository) Update(id uint, cupon *Cupon) error {
	return r.db.Model(&Cupon{}).Where("id_cupon = ? AND estado = ?", id, true).Updates(cupon).Error
}

func (r *cuponRepository) Delete(id uint) error {
	return r.db.Model(&Cupon{}).Where("id_cupon = ? AND estado = ?", id, true).Update("estado", false).Error
}

func (r *cuponRepository) GetByCodigo(codigo string) (*Cupon, error) {
	var cupon Cupon
	err := r.db.Where("codigo = ? AND estado = ?", codigo, true).First(&cupon).Error
	if err != nil {
		return nil, err
	}
	return &cupon, nil
}
