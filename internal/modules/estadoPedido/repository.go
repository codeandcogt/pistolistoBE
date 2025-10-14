package estadoPedido

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type EstadoPedidoRepository interface {
	Create(estado *EstadoPedido) error
	GetByID(id uint) (*EstadoPedido, error)
	GetAll() ([]*EstadoPedido, error)
	Update(estado *EstadoPedido) (string, error)
	Delete(id uint) (string, error)
}

type estadoPedidoRepository struct {
	db *gorm.DB
}

func NewEstadoPedidoRepository(db *gorm.DB) EstadoPedidoRepository {
	return &estadoPedidoRepository{db}
}

func (r *estadoPedidoRepository) Create(estado *EstadoPedido) error {
	return r.db.Create(estado).Error
}

func (r *estadoPedidoRepository) GetByID(id uint) (*EstadoPedido, error) {
	var estado EstadoPedido
	err := r.db.Where("estado = ?", true).First(&estado, id).Error
	if err != nil {
		return nil, err
	}
	return &estado, nil
}

func (r *estadoPedidoRepository) GetAll() ([]*EstadoPedido, error) {
	var estados []*EstadoPedido
	err := r.db.Where("estado = ?", true).Find(&estados).Error
	if err != nil {
		return nil, err
	}
	return estados, nil
}

func (r *estadoPedidoRepository) Update(estado *EstadoPedido) (string, error) {
	result := r.db.Model(&EstadoPedido{}).
		Where("id_estado_pedido = ? AND estado = ?", estado.IdEstadoPedido, true).
		Updates(estado)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_UPDATED, nil
}

func (r *estadoPedidoRepository) Delete(id uint) (string, error) {
	result := r.db.Model(&EstadoPedido{}).
		Where("id_estado_pedido = ?", id).
		Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}
