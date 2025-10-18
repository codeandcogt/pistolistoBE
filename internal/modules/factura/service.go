package factura

import (
	"errors"
	"fmt"
	"math/rand"
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
	repo FacturaRepository
}

func NewFacturaService(repo FacturaRepository) FacturaService {
	return &facturaService{repo}
}

func generarNumeroSerie() (string, string) {
	serie := "A"
	numero := fmt.Sprintf("%06d", rand.Intn(999999))
	return numero, serie
}

func generarUUID() string {
	return fmt.Sprintf("UUID-%d-%d", time.Now().Unix(), rand.Intn(99999))
}

func (s *facturaService) EmitirFactura(factura *Factura) (*Factura, error) {
	if factura.IdPedido == 0 {
		return nil, errors.New("pedido no válido para facturación")
	}

	numero, serie := generarNumeroSerie()
	factura.Numero = numero
	factura.Serie = serie
	factura.UUID = generarUUID()
	factura.FechaEmision = time.Now()

	if factura.Total == 0 {
		factura.Total = factura.Subtotal + factura.Impuestos - factura.Descuento
	}

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
