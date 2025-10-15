package pedido

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type PedidoRepository interface {
	Create(pedido *Pedido) error
	GetByID(id uint) (*Pedido, error)
	GetAll() ([]*Pedido, error)
	GetByCliente(idCliente uint) ([]*Pedido, error)
	Update(pedido *Pedido) (string, error)
	UpdateEstado(idPedido uint, idEstadoPedido int) error
	Delete(id uint) (string, error)
}

type pedidoRepository struct {
	db *gorm.DB
}

func NewPedidoRepository(db *gorm.DB) PedidoRepository {
	return &pedidoRepository{db}
}

func (r *pedidoRepository) Create(pedido *Pedido) error {
	return r.db.Create(pedido).Error
}

func (r *pedidoRepository) GetByID(id uint) (*Pedido, error) {
	var pedido Pedido
	err := r.db.Where("estado = ?", true).First(&pedido, id).Error
	if err != nil {
		return nil, err
	}
	return &pedido, nil
}

func (r *pedidoRepository) GetAll() ([]*Pedido, error) {
	var pedidos []*Pedido
	err := r.db.Where("estado = ?", true).Find(&pedidos).Error
	if err != nil {
		return nil, err
	}
	return pedidos, nil
}

func (r *pedidoRepository) GetByCliente(idCliente uint) ([]*Pedido, error) {
	var pedidos []*Pedido
	err := r.db.Where("estado = ? AND id_cliente = ?", true, idCliente).Find(&pedidos).Error
	if err != nil {
		return nil, err
	}
	return pedidos, nil
}

func (r *pedidoRepository) Update(pedido *Pedido) (string, error) {
	result := r.db.Model(&Pedido{}).
		Where("id_pedido = ? AND estado = ?", pedido.IdPedido, true).
		Updates(pedido)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_UPDATED, nil
}

func (r *pedidoRepository) UpdateEstado(idPedido uint, idEstadoPedido int) error {
	return r.db.Model(&Pedido{}).
		Where("id_pedido = ? AND estado = ?", idPedido, true).
		Update("id_estado_pedido", idEstadoPedido).Error
}

func (r *pedidoRepository) Delete(id uint) (string, error) {
	result := r.db.Model(&Pedido{}).Where("id_pedido = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}
