package cuota

import "gorm.io/gorm"

type CuotaRepository interface {
	Create(cuota *Cuota) error
	GetByID(id uint) (*Cuota, error)
	GetAll() ([]*Cuota, error)
	Update(id uint, cuota *Cuota) error
	Delete(id uint) error
}

type cuotaRepository struct {
	db *gorm.DB
}

func NewCuotaRepository(db *gorm.DB) CuotaRepository {
	return &cuotaRepository{db}
}

func (r *cuotaRepository) Create(cuota *Cuota) error {
	return r.db.Create(cuota).Error
}

func (r *cuotaRepository) GetByID(id uint) (*Cuota, error) {
	var cuota Cuota
	err := r.db.Where("id_cuota = ? AND estado = ?", id, true).First(&cuota).Error
	if err != nil {
		return nil, err
	}
	return &cuota, nil
}

func (r *cuotaRepository) GetAll() ([]*Cuota, error) {
	var cuotas []*Cuota
	err := r.db.Where("estado = ?", true).Order("fecha_creacion DESC").Find(&cuotas).Error
	if err != nil {
		return nil, err
	}
	return cuotas, nil
}

func (r *cuotaRepository) Update(id uint, cuota *Cuota) error {
	return r.db.Model(&Cuota{}).Where("id_cuota = ? AND estado = ?", id, true).Updates(cuota).Error
}

func (r *cuotaRepository) Delete(id uint) error {
	return r.db.Model(&Cuota{}).Where("id_cuota = ? AND estado = ?", id, true).Update("estado", false).Error
}
