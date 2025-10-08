package municipality

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type MunicipalityRepository interface {
	Create(municipality *Municipality) error
	GetByID(id uint) (*Municipality, error)
	GetAll() ([]*Municipality, error)
	Update(municipality *Municipality) (string, error)
	Delete(id uint) (string, error)
}

type municipalityRepository struct {
	db *gorm.DB
}

func NewMunicipalityRepository(db *gorm.DB) MunicipalityRepository {
	return &municipalityRepository{db}
}

func (r *municipalityRepository) Create(municipality *Municipality) error {
	return r.db.Create(municipality).Error
}

func (r *municipalityRepository) GetByID(id uint) (*Municipality, error) {
	var municipality Municipality
	err := r.db.Where("estado = ?", true).First(&municipality, id).Error
	if err != nil {
		return nil, err
	}
	return &municipality, nil
}

func (r *municipalityRepository) GetAll() ([]*Municipality, error) {
	var municipalities []*Municipality
	err := r.db.Where("estado = ?", true).Find(&municipalities).Error
	if err != nil {
		return nil, err
	}
	return municipalities, nil
}

func (r *municipalityRepository) Update(municipality *Municipality) (string, error) {
	result := r.db.Model(&Municipality{}).
		Where("id_municipio = ? AND estado = ?", municipality.IdMunicipio, true).
		Updates(municipality)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_UPDATED, nil
}

func (r *municipalityRepository) Delete(id uint) (string, error) {
	result := r.db.Model(&Municipality{}).Where("id_municipio = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}
	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}
	return common.SUCCESS_DELETED, nil
}
