package factura

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type FacturaRepository interface {
	Create(factura *Factura) error
	GetByID(id uint) (*Factura, error)
	GetAll() ([]*Factura, error)
	GetByPedido(idPedido uint) (*Factura, error)
	Update(factura *Factura) (string, error)
	Delete(id uint) (string, error)
}

type facturaRepository struct {
	db *gorm.DB
}

func NewFacturaRepository(db *gorm.DB) FacturaRepository {
	return &facturaRepository{db}
}

func (r *facturaRepository) Create(factura *Factura) error {
	return r.db.Create(factura).Error
}

func (r *facturaRepository) GetByID(id uint) (*Factura, error) {
	var factura Factura
	err := r.db.Where("estado = ?", true).First(&factura, id).Error
	if err != nil {
		return nil, err
	}
	return &factura, nil
}

func (r *facturaRepository) GetAll() ([]*Factura, error) {
	var facturas []*Factura
	err := r.db.Where("estado = ?", true).Find(&facturas).Error
	if err != nil {
		return nil, err
	}
	return facturas, nil
}

func (r *facturaRepository) GetByPedido(idPedido uint) (*Factura, error) {
	var factura Factura
	err := r.db.Where("estado = ? AND id_pedido = ?", true, idPedido).First(&factura).Error
	if err != nil {
		return nil, err
	}
	return &factura, nil
}

func (r *facturaRepository) Update(factura *Factura) (string, error) {
	result := r.db.Model(&Factura{}).
		Where("id_factura = ? AND estado = ?", factura.IdFactura, true).
		Updates(factura)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}
	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}
	return common.SUCCESS_UPDATED, nil
}

func (r *facturaRepository) Delete(id uint) (string, error) {
	result := r.db.Model(&Factura{}).Where("id_factura = ?", id).Update("estado", false)
	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}
	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}
	return common.SUCCESS_DELETED, nil
}
