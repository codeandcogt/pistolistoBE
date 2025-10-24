package logUbicacion

import (
	"gorm.io/gorm"
)

type LogUbicacionRepository interface {
	Create(log *LogUbicacion) error
	GetByPiloto(pilotoId uint) ([]*LogUbicacion, error)
	GetLastByPiloto(pilotoId uint) (*LogUbicacion, error)
}

type logUbicacionRepository struct {
	db *gorm.DB
}

func NewLogUbicacionRepository(db *gorm.DB) LogUbicacionRepository {
	return &logUbicacionRepository{db}
}

func (r *logUbicacionRepository) Create(log *LogUbicacion) error {
	return r.db.Create(log).Error
}

func (r *logUbicacionRepository) GetByPiloto(pilotoId uint) ([]*LogUbicacion, error) {
	var logs []*LogUbicacion
	err := r.db.Where("estado = ? AND piloto_id = ?", true, pilotoId).
		Order("fecha_creacion DESC").Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (r *logUbicacionRepository) GetLastByPiloto(pilotoId uint) (*LogUbicacion, error) {
	var log LogUbicacion
	err := r.db.Where("estado = ? AND piloto_id = ?", true, pilotoId).
		Order("fecha_creacion DESC").First(&log).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}
