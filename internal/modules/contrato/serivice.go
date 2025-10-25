package contrato

import (
	"fmt"
	"time"
)

type ContratoService interface {
	CrearContratoAutomatico(idAvaluo uint, monto float64) (*Contrato, error)
	GetByAvaluo(idAvaluo uint) (*Contrato, error)
	GetAll() ([]*Contrato, error)
}

type contratoService struct {
	repo ContratoRepository
}

func NewContratoService(repo ContratoRepository) ContratoService {
	return &contratoService{repo}
}

// Genera contrato automáticamente según el monto avaluado
func (s *contratoService) CrearContratoAutomatico(idAvaluo uint, monto float64) (*Contrato, error) {
	tipo := "Contrato Estándar"
	var condiciones string

	switch {
	case monto < 5000:
		condiciones = "Condiciones básicas, sin requerir garantía adicional."
	case monto >= 5000 && monto < 20000:
		tipo = "Contrato Preferente"
		condiciones = "Incluye revisión semestral y garantía mínima."
	default:
		tipo = "Contrato Premium"
		condiciones = "Garantía extendida y mantenimiento anual incluido."
	}

	contrato := &Contrato{
		IdAvaluo:              int(idAvaluo),
		TipoContrato:          tipo,
		CondicionesEspeciales: &condiciones,
		FechaCreacion:         new(time.Time),
	}

	now := time.Now()
	*contrato.FechaCreacion = now

	if err := s.repo.Create(contrato); err != nil {
		return nil, err
	}

	fmt.Printf("✅ Contrato generado automáticamente para Avalúo #%d\n", idAvaluo)
	return contrato, nil
}

func (s *contratoService) GetByAvaluo(idAvaluo uint) (*Contrato, error) {
	return s.repo.GetByAvaluo(idAvaluo)
}

func (s *contratoService) GetAll() ([]*Contrato, error) {
	return s.repo.GetAll()
}
