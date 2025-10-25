package avaluo

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type AvaluoRepository interface {
	Create(avaluo *Avaluo) error
	GetByID(id uint) (*Avaluo, error)
	GetAll() ([]*Avaluo, error)
	GetByFormulario(idFormulario uint) (*Avaluo, error)
	Update(avaluo *Avaluo) (string, error)
	Delete(id uint) (string, error)
}

type avaluoRepository struct {
	db *gorm.DB
}

func NewAvaluoRepository(db *gorm.DB) AvaluoRepository {
	return &avaluoRepository{db}
}

func (r *avaluoRepository) Create(avaluo *Avaluo) error {
	return r.db.Create(avaluo).Error
}

func (r *avaluoRepository) GetByID(id uint) (*Avaluo, error) {
	var avaluo Avaluo
	err := r.db.Where("estado = ?", true).First(&avaluo, id).Error
	if err != nil {
		return nil, err
	}
	return &avaluo, nil
}

func (r *avaluoRepository) GetAll() ([]*Avaluo, error) {
	var avaluos []*Avaluo
	err := r.db.Where("estado = ?", true).Find(&avaluos).Error
	if err != nil {
		return nil, err
	}
	return avaluos, nil
}

func (r *avaluoRepository) GetByFormulario(idFormulario uint) (*Avaluo, error) {
	var avaluo Avaluo
	err := r.db.Where("estado = ? AND id_formulario = ?", true, idFormulario).First(&avaluo).Error
	if err != nil {
		return nil, err
	}
	return &avaluo, nil
}

func (r *avaluoRepository) Update(avaluo *Avaluo) (string, error) {
	result := r.db.Model(&Avaluo{}).Where("id_avaluo = ? AND estado = ?", avaluo.IdAvaluo, true).Updates(avaluo)
	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}
	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}
	return common.SUCCESS_UPDATED, nil
}

func (r *avaluoRepository) Delete(id uint) (string, error) {
	result := r.db.Model(&Avaluo{}).Where("id_avaluo = ?", id).Update("estado", false)
	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}
	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}
	return common.SUCCESS_DELETED, nil
}
