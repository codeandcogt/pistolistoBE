package logUbicacionTiempoReal

import "gorm.io/gorm"

type LogUbicacionTiempoRealRepository interface {
	Upsert(pilotoId int, lat, lon float64) error
	GetByPiloto(pilotoId uint) (*LogUbicacionTiempoReal, error)
	GetAll() ([]*LogUbicacionTiempoReal, error)
}

type logUbicacionTiempoRealRepository struct {
	db *gorm.DB
}

func NewLogUbicacionTiempoRealRepository(db *gorm.DB) LogUbicacionTiempoRealRepository {
	return &logUbicacionTiempoRealRepository{db}
}

func (r *logUbicacionTiempoRealRepository) Upsert(pilotoId int, lat, lon float64) error {
	var log LogUbicacionTiempoReal
	err := r.db.Where("piloto_id = ?", pilotoId).First(&log).Error
	if err == gorm.ErrRecordNotFound {
		log = LogUbicacionTiempoReal{
			PilotoId: pilotoId,
			Latitud:  lat,
			Longitud: lon,
		}
		return r.db.Create(&log).Error
	}

	return r.db.Model(&log).Updates(map[string]interface{}{
		"latitud":  lat,
		"longitud": lon,
	}).Error
}

func (r *logUbicacionTiempoRealRepository) GetByPiloto(pilotoId uint) (*LogUbicacionTiempoReal, error) {
	var log LogUbicacionTiempoReal
	if err := r.db.Where("piloto_id = ?", pilotoId).First(&log).Error; err != nil {
		return nil, err
	}
	return &log, nil
}

func (r *logUbicacionTiempoRealRepository) GetAll() ([]*LogUbicacionTiempoReal, error) {
	var logs []*LogUbicacionTiempoReal
	if err := r.db.Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}
