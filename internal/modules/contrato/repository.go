package contrato

import "gorm.io/gorm"

type ContratoRepository interface {
	Create(contrato *Contrato) error
	GetByAvaluo(idAvaluo uint) (*Contrato, error)
	GetAll() ([]*Contrato, error)
	GetByID(id uint) (*Contrato, error)
}

type contratoRepository struct {
	db *gorm.DB
}

func NewContratoRepository(db *gorm.DB) ContratoRepository {
	return &contratoRepository{db}
}

func (r *contratoRepository) Create(contrato *Contrato) error {
	return r.db.Create(contrato).Error
}

func (r *contratoRepository) GetByAvaluo(idAvaluo uint) (*Contrato, error) {
	var contrato Contrato
	err := r.db.Where("id_avaluo = ? AND estado = ?", idAvaluo, true).First(&contrato).Error
	if err != nil {
		return nil, err
	}
	return &contrato, nil
}

func (r *contratoRepository) GetByID(id uint) (*Contrato, error) {
	var contrato Contrato
	err := r.db.Where("id_contrato = ? AND estado = ?", id, true).First(&contrato).Error
	if err != nil {
		return nil, err
	}
	return &contrato, nil
}

func (r *contratoRepository) GetAll() ([]*Contrato, error) {
	var contratos []*Contrato
	err := r.db.Where("estado = ?", true).Find(&contratos).Error
	if err != nil {
		return nil, err
	}
	return contratos, nil
}
