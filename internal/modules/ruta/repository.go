package ruta

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type RutaRepository interface {
	Create(ruta *Ruta) error
	GetByID(id uint) (*Ruta, error)
	GetAll() ([]*Ruta, error)
	Update(ruta *Ruta) (string, error)
	Delete(id uint) (string, error)
	UpdateEstado(idRuta uint, idEstadoRuta int) error
}

type rutaRepository struct {
	db *gorm.DB
}

func NewRutaRepository(db *gorm.DB) RutaRepository {
	return &rutaRepository{db}
}

func (r *rutaRepository) Create(ruta *Ruta) error {
	return r.db.Create(ruta).Error
}

func (r *rutaRepository) GetByID(id uint) (*Ruta, error) {
	var ruta Ruta
	err := r.db.Where("estado = ?", true).First(&ruta, id).Error
	if err != nil {
		return nil, err
	}
	return &ruta, nil
}

func (r *rutaRepository) GetAll() ([]*Ruta, error) {
	var rutas []*Ruta
	err := r.db.Where("estado = ?", true).Find(&rutas).Error
	if err != nil {
		return nil, err
	}
	return rutas, nil
}

func (r *rutaRepository) Update(ruta *Ruta) (string, error) {
	result := r.db.Model(&Ruta{}).
		Where("id_ruta = ? AND estado = ?", ruta.IdRuta, true).
		Updates(ruta)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}
	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}
	return common.SUCCESS_UPDATED, nil
}

func (r *rutaRepository) Delete(id uint) (string, error) {
	result := r.db.Model(&Ruta{}).Where("id_ruta = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}
	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}
	return common.SUCCESS_DELETED, nil
}

func (r *rutaRepository) UpdateEstado(idRuta uint, idEstadoRuta int) error {
	return r.db.Model(&Ruta{}).
		Where("id_ruta = ? AND estado = ?", idRuta, true).
		Update("id_estado_ruta", idEstadoRuta).Error
}
