package vehiculo

import (
	"gorm.io/gorm"
)

type VehiculoRepository interface {
	Create(vehiculo *Vehiculo) error
	GetByID(id uint) (*Vehiculo, error)
	GetAll() ([]*Vehiculo, error)
	GetByPilotoID(pilotoId uint) ([]*Vehiculo, error)
	Update(id uint, vehiculo *Vehiculo) error
	Delete(id uint) error
}

type vehiculoRepository struct {
	db *gorm.DB
}

func NewVehiculoRepository(db *gorm.DB) VehiculoRepository {
	return &vehiculoRepository{db}
}

func (r *vehiculoRepository) Create(vehiculo *Vehiculo) error {
	return r.db.Create(vehiculo).Error
}

func (r *vehiculoRepository) GetByID(id uint) (*Vehiculo, error) {
	var vehiculo Vehiculo
	err := r.db.Where("id_vehiculo = ? AND estado = ?", id, true).First(&vehiculo).Error
	if err != nil {
		return nil, err
	}
	return &vehiculo, nil
}

func (r *vehiculoRepository) GetAll() ([]*Vehiculo, error) {
	var vehiculos []*Vehiculo
	err := r.db.Where("estado = ?", true).Order("fecha_creacion DESC").Find(&vehiculos).Error
	if err != nil {
		return nil, err
	}
	return vehiculos, nil
}

func (r *vehiculoRepository) GetByPilotoID(pilotoId uint) ([]*Vehiculo, error) {
	var vehiculos []*Vehiculo
	err := r.db.Where("piloto_id = ? AND estado = ?", pilotoId, true).Find(&vehiculos).Error
	if err != nil {
		return nil, err
	}
	return vehiculos, nil
}

func (r *vehiculoRepository) Update(id uint, vehiculo *Vehiculo) error {
	return r.db.Model(&Vehiculo{}).Where("id_vehiculo = ? AND estado = ?", id, true).Updates(vehiculo).Error
}

func (r *vehiculoRepository) Delete(id uint) error {
	return r.db.Model(&Vehiculo{}).Where("id_vehiculo = ? AND estado = ?", id, true).Update("estado", false).Error
}
