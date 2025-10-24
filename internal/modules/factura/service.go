package factura

import (
	"errors"
	"fmt"
	"math/rand"
	"pistolistoBE/internal/modules/pedido"
	"time"
)

type FacturaService interface {
	EmitirFactura(factura *Factura) (*Factura, error)
	GetByID(id uint) (*Factura, error)
	GetAll() ([]*Factura, error)
	GetByPedido(idPedido uint) (*Factura, error)
	Update(factura *Factura) (string, error)
	Delete(id uint) (string, error)
}

type facturaService struct {
	repo       FacturaRepository
	pedidoRepo pedido.PedidoRepository
}

func NewFacturaService(repo FacturaRepository, pedidoRepo pedido.PedidoRepository) FacturaService {
	return &facturaService{repo, pedidoRepo}
}

// Generadores de datos de factura
func generarNumeroSerie() (string, string) {
	serie := "A"
	numero := fmt.Sprintf("%06d", rand.Intn(999999))
	return numero, serie
}

func generarUUID() string {
	return fmt.Sprintf("UUID-%d-%d", time.Now().Unix(), rand.Intn(99999))
}

func (s *facturaService) EmitirFactura(factura *Factura) (*Factura, error) {
	// Validar pedido existente
	pedidoData, err := s.pedidoRepo.GetByID(uint(factura.IdPedido))
	if err != nil {
		return nil, fmt.Errorf("no se encontró el pedido con id %d", factura.IdPedido)
	}

	// Validar que el pedido esté pagado
	if pedidoData.IdEstadoPedido != 2 {
		return nil, fmt.Errorf("el pedido %d aún no está pagado, no se puede emitir factura", pedidoData.IdPedido)
	}

	// Verificar si ya existe factura para este pedido
	existente, _ := s.repo.GetByPedido(uint(factura.IdPedido))
	if existente != nil && existente.IdFactura != 0 {
		return nil, fmt.Errorf("ya existe una factura asociada al pedido %d", factura.IdPedido)
	}

	// Completar datos financieros desde el pedido
	factura.Subtotal = pedidoData.Subtotal
	factura.Impuestos = pedidoData.Impuesto
	factura.Total = pedidoData.Total
	factura.Descuento = 0
	factura.FechaEmision = time.Now()

	// Generar número, serie y UUID
	numero, serie := generarNumeroSerie()
	factura.Numero = numero
	factura.Serie = serie
	factura.UUID = generarUUID()

	// Validar campos obligatorios
	if factura.RazonSocial == "" || factura.DireccionFiscal == "" {
		return nil, errors.New("la razón social y la dirección fiscal son obligatorias")
	}

	// Guardar factura
	if err := s.repo.Create(factura); err != nil {
		return nil, err
	}

	return factura, nil
}

func (s *facturaService) GetByID(id uint) (*Factura, error) {
	return s.repo.GetByID(id)
}

func (s *facturaService) GetAll() ([]*Factura, error) {
	return s.repo.GetAll()
}

func (s *facturaService) GetByPedido(idPedido uint) (*Factura, error) {
	return s.repo.GetByPedido(idPedido)
}

func (s *facturaService) Update(factura *Factura) (string, error) {
	return s.repo.Update(factura)
}

func (s *facturaService) Delete(id uint) (string, error) {
	return s.repo.Delete(id)
}
