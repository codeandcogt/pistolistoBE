package pago

import (
	"errors"
	"fmt"
	"math/rand"
	"pistolistoBE/internal/modules/pedido"
	"time"
)

type PagoService interface {
	ProcesarPago(pago *Pago) (*Pago, error)
	GetByID(id uint) (*Pago, error)
	GetAll() ([]*Pago, error)
	GetByPedido(idPedido uint) ([]*Pago, error)
	Update(pago *Pago) (string, error)
	Delete(id uint) (string, error)
}

type pagoService struct {
	repo       PagoRepository
	pedidoRepo pedido.PedidoRepository
}

func NewPagoService(repo PagoRepository, pedidoRepo pedido.PedidoRepository) PagoService {
	return &pagoService{repo, pedidoRepo}
}

// Mock pasarela de pago
func (s *pagoService) procesarMockPasarela(pago *Pago) (string, error) {
	time.Sleep(400 * time.Millisecond)
	success := rand.Intn(100) < 95
	if success {
		return fmt.Sprintf("AUTH-%d", rand.Intn(999999)), nil
	}
	return "", errors.New("transacción rechazada por la pasarela")
}

func (s *pagoService) ProcesarPago(pago *Pago) (*Pago, error) {
	if pago.Monto <= 0 {
		return nil, errors.New("el monto del pago debe ser mayor que cero")
	}

	auth, err := s.procesarMockPasarela(pago)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	pago.Autorizacion = &auth
	pago.FechaAutorizacion = &now

	if err := s.repo.Create(pago); err != nil {
		return nil, err
	}

	if err := s.pedidoRepo.UpdateEstado(uint(pago.IdPedido), 2); err != nil {
		fmt.Println("No se pudo actualizar el estado del pedido:", err)
	}

	return pago, nil
}

func (s *pagoService) GetByID(id uint) (*Pago, error) {
	return s.repo.GetByID(id)
}

func (s *pagoService) GetAll() ([]*Pago, error) {
	return s.repo.GetAll()
}

func (s *pagoService) GetByPedido(idPedido uint) ([]*Pago, error) {
	return s.repo.GetByPedido(idPedido)
}

func (s *pagoService) Update(pago *Pago) (string, error) {
	if pago.IdPago == 0 {
		return "", errors.New("id_pago es requerido")
	}
	return s.repo.Update(pago)
}

func (s *pagoService) Delete(id uint) (string, error) {
	return s.repo.Delete(id)
}
