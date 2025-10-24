package estadoRuta

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type EstadoRutaRepository interface {
	Create(estado *EstadoRuta) error
	GetByID(id uint) (*EstadoRuta, error)
	GetAll() ([]*EstadoRuta, error)
	Update(estado *EstadoRuta) (string, error)
	Delete(id uint) (string, error)
}

type estadoRutaRepository struct {
	db *gorm.DB
}

func NewEstadoRutaRepository(db *gorm.DB) EstadoRutaRepository {
	return &estadoRutaRepository{db}
}

func (r *estadoRutaRepository) Create(estado *EstadoRuta) error {
	return r.db.Create(estado).Error
}

func (r *estadoRutaRepository) GetByID(id uint) (*EstadoRuta, error) {
	var estado EstadoRuta
	err := r.db.Where("estado = ?", true).First(&estado, id).Error
	if err != nil {
		return nil, err
	}
	return &estado, nil
}

func (r *estadoRutaRepository) GetAll() ([]*EstadoRuta, error) {
	var estados []*EstadoRuta
	err := r.db.Where("estado = ?", true).Find(&estados).Error
	if err != nil {
		return nil, err
	}
	return estados, nil
}

func (r *estadoRutaRepository) Update(estado *EstadoRuta) (string, error) {
	result := r.db.Model(&EstadoRuta{}).
		Where("id_estado_ruta = ? AND estado = ?", estado.IdEstadoRuta, true).
		Updates(estado)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}
	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}
	return common.SUCCESS_UPDATED, nil
}

func (r *estadoRutaRepository) Delete(id uint) (string, error) {
	result := r.db.Model(&EstadoRuta{}).
		Where("id_estado_ruta = ?", id).
		Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}
	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}
	return common.SUCCESS_DELETED, nil
}
