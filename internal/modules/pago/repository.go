package pago

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type PagoRepository interface {
	Create(pago *Pago) error
	GetByID(id uint) (*Pago, error)
	GetAll() ([]*Pago, error)
	GetByPedido(idPedido uint) ([]*Pago, error)
	Update(pago *Pago) (string, error)
	Delete(id uint) (string, error)
}

type pagoRepository struct {
	db *gorm.DB
}

func NewPagoRepository(db *gorm.DB) PagoRepository {
	return &pagoRepository{db}
}

func (r *pagoRepository) Create(pago *Pago) error {
	return r.db.Create(pago).Error
}

func (r *pagoRepository) GetByID(id uint) (*Pago, error) {
	var pago Pago
	err := r.db.Where("estado = ?", true).First(&pago, id).Error
	if err != nil {
		return nil, err
	}
	return &pago, nil
}

func (r *pagoRepository) GetAll() ([]*Pago, error) {
	var pagos []*Pago
	err := r.db.Where("estado = ?", true).Find(&pagos).Error
	if err != nil {
		return nil, err
	}
	return pagos, nil
}

func (r *pagoRepository) GetByPedido(idPedido uint) ([]*Pago, error) {
	var pagos []*Pago
	err := r.db.Where("estado = ? AND id_pedido = ?", true, idPedido).Find(&pagos).Error
	if err != nil {
		return nil, err
	}
	return pagos, nil
}

func (r *pagoRepository) Update(pago *Pago) (string, error) {
	result := r.db.Model(&Pago{}).
		Where("id_pago = ? AND estado = ?", pago.IdPago, true).
		Updates(pago)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_UPDATED, nil
}

func (r *pagoRepository) Delete(id uint) (string, error) {
	result := r.db.Model(&Pago{}).Where("id_pago = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}
