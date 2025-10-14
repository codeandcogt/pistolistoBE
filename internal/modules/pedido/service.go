package pedido

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type PedidoService interface {
	Checkout(pedido *Pedido) (*Pedido, error)
	GetByID(id uint) (*Pedido, error)
	GetAll() ([]*Pedido, error)
	GetByCliente(idCliente uint) ([]*Pedido, error)
	CambiarEstado(idPedido uint, idEstadoPedido int) error
	CancelarPedido(idPedido uint, motivo string) error
	Delete(id uint) (string, error)
}

type pedidoService struct {
	repo PedidoRepository
}

func NewPedidoService(repo PedidoRepository) PedidoService {
	return &pedidoService{repo}
}

// ----------------------
// Crear pedido (checkout)
// ----------------------
func (s *pedidoService) Checkout(pedido *Pedido) (*Pedido, error) {
	// Generar número de pedido único
	pedido.NumeroPedido = fmt.Sprintf("PED-%d-%04d", time.Now().Unix(), rand.Intn(10000))

	// Simular cálculo de totales (en un sistema real, esto vendría del carrito)
	if pedido.Subtotal == 0 {
		return nil, errors.New("el subtotal no puede ser cero")
	}
	pedido.Impuesto = pedido.Subtotal * 0.12 // 12% IVA ejemplo
	pedido.Total = pedido.Subtotal + pedido.Impuesto + pedido.CostoEnvio

	// Estado inicial “Pendiente”
	pedido.IdEstadoPedido = 1

	if err := s.repo.Create(pedido); err != nil {
		return nil, err
	}
	return pedido, nil
}

func (s *pedidoService) GetByID(id uint) (*Pedido, error) {
	return s.repo.GetByID(id)
}

func (s *pedidoService) GetAll() ([]*Pedido, error) {
	return s.repo.GetAll()
}

func (s *pedidoService) GetByCliente(idCliente uint) ([]*Pedido, error) {
	return s.repo.GetByCliente(idCliente)
}

func (s *pedidoService) CambiarEstado(idPedido uint, idEstadoPedido int) error {
	return s.repo.UpdateEstado(idPedido, idEstadoPedido)
}

func (s *pedidoService) CancelarPedido(idPedido uint, motivo string) error {
	// En una implementación real se validaría si el pedido aún no fue enviado
	return s.repo.UpdateEstado(idPedido, 6) // Estado “Cancelado”
}

func (s *pedidoService) Delete(id uint) (string, error) {
	return s.repo.Delete(id)
}
