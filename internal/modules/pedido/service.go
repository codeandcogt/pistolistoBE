package pedido

import (
	"errors"
	"fmt"
	"math/rand"
	"pistolistoBE/internal/modules/carrito"
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
	repo        PedidoRepository
	carritoRepo carrito.CarritoRepository
}

func NewPedidoService(repo PedidoRepository, carritoRepo carrito.CarritoRepository) PedidoService {
	return &pedidoService{repo: repo, carritoRepo: carritoRepo}
}

// ----------------------
// Crear pedido (checkout)
// ----------------------
func (s *pedidoService) Checkout(pedido *Pedido) (*Pedido, error) {
	// 1️ Validar carrito existente
	cart, err := s.carritoRepo.GetByID(uint(pedido.IdCarrito))
	if err != nil {
		return nil, errors.New("carrito no encontrado")
	}
	if len(cart.Items) == 0 {
		return nil, errors.New("el carrito está vacío")
	}

	// 2 Calcular totales
	var subtotal float64
	for _, item := range cart.Items {
		if item.Subtotal != nil {
			subtotal += *item.Subtotal
		}
	}

	descuento := 0.0
	if cart.Descuento != nil {
		descuento = *cart.Descuento
	}
	impuesto := (subtotal - descuento) * 0.12 // IVA 12%
	costoEnvio := 50.00                       // Temporal: fijo o según dirección
	total := subtotal - descuento + impuesto + costoEnvio

	// Generar datos del pedido tomando cliente desde el carrito
	if cart.ClienteId == nil {
		return nil, errors.New("el carrito no tiene cliente asociado")
	}

	pedido.IdCliente = int(*cart.ClienteId)
	pedido.Subtotal = subtotal
	pedido.Impuesto = impuesto
	pedido.CostoEnvio = costoEnvio
	pedido.Total = total
	pedido.IdEstadoPedido = 1 // Pendiente
	pedido.NumeroPedido = fmt.Sprintf("PED-%d-%04d", time.Now().Unix(), rand.Intn(10000))

	// Guardar pedido
	if err := s.repo.Create(pedido); err != nil {
		return nil, err
	}

	//  Enlazar carrito con pedido creado
	cart.PedidoId = &pedido.IdPedido
	err = s.carritoRepo.Update(cart.IdCarrito, cart)
	if err != nil {
		fmt.Println("⚠️ No se pudo actualizar el carrito con el pedido:", err)
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
	return s.repo.UpdateEstado(idPedido, 6) // Estado “Cancelado”
}

func (s *pedidoService) Delete(id uint) (string, error) {
	return s.repo.Delete(id)
}
