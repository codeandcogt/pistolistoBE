package piloto

import (
	"gorm.io/gorm"
)

type PilotoRepository interface {
	Create(piloto *Piloto) error
	GetByID(id uint) (*Piloto, error)
	GetAll() ([]*Piloto, error)
	GetByAdministrativoID(administrativoId uint) ([]*Piloto, error)
	Update(id uint, piloto *Piloto) error
	Delete(id uint) error
}

type pilotoRepository struct {
	db *gorm.DB
}

func NewPilotoRepository(db *gorm.DB) PilotoRepository {
	return &pilotoRepository{db}
}

func (r *pilotoRepository) Create(piloto *Piloto) error {
	return r.db.Create(piloto).Error
}

func (r *pilotoRepository) GetByID(id uint) (*Piloto, error) {
	var piloto Piloto
	err := r.db.Where("id_piloto = ? AND estado = ?", id, true).First(&piloto).Error
	if err != nil {
		return nil, err
	}
	return &piloto, nil
}

func (r *pilotoRepository) GetAll() ([]*Piloto, error) {
	var pilotos []*Piloto
	err := r.db.Where("estado = ?", true).Order("fecha_creacion DESC").Find(&pilotos).Error
	if err != nil {
		return nil, err
	}
	return pilotos, nil
}

func (r *pilotoRepository) GetByAdministrativoID(administrativoId uint) ([]*Piloto, error) {
	var pilotos []*Piloto
	err := r.db.Where("administrativo_id = ? AND estado = ?", administrativoId, true).Find(&pilotos).Error
	if err != nil {
		return nil, err
	}
	return pilotos, nil
}

func (r *pilotoRepository) Update(id uint, piloto *Piloto) error {
	return r.db.Model(&Piloto{}).Where("id_piloto = ? AND estado = ?", id, true).Updates(piloto).Error
}

func (r *pilotoRepository) Delete(id uint) error {
	return r.db.Model(&Piloto{}).Where("id_piloto = ? AND estado = ?", id, true).Update("estado", false).Error
}
