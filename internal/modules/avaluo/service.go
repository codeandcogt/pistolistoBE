package avaluo

import (
	"errors"
	"math/rand"
	"pistolistoBE/internal/modules/formulario"
	"time"
)

type AvaluoService interface {
	CrearAvaluoManual(avaluo *Avaluo) error
	CrearAvaluoAutomatico(idFormulario uint) (*Avaluo, error)
	GetByID(id uint) (*Avaluo, error)
	GetAll() ([]*Avaluo, error)
	Delete(id uint) (string, error)
}

type avaluoService struct {
	repo     AvaluoRepository
	formRepo formulario.FormularioRepository
}

func NewAvaluoService(repo AvaluoRepository, formRepo formulario.FormularioRepository) AvaluoService {
	return &avaluoService{repo, formRepo}
}

// Avalúo manual (usuario lo crea explícitamente)
func (s *avaluoService) CrearAvaluoManual(avaluo *Avaluo) error {
	avaluo.FechaAvaluo = time.Now()
	return s.repo.Create(avaluo)
}

// Avalúo automático basado en los datos del formulario
func (s *avaluoService) CrearAvaluoAutomatico(idFormulario uint) (*Avaluo, error) {
	form, err := s.formRepo.GetByID(idFormulario)
	if err != nil {
		return nil, errors.New("no se encontró el formulario")
	}

	// Simulación: Algoritmo de evaluación de avalúo
	// Basado en MontoSolicitado, tipo de operación, y aleatoriedad
	var factor float64 = 1.0
	if form.TipoOperacion == "compra" {
		factor = 1.05
	} else if form.TipoOperacion == "venta" {
		factor = 0.95
	}
	if rand.Float64() > 0.7 {
		factor += 0.03
	}

	montoAvaluado := form.MontoSolicitado * factor

	avaluo := &Avaluo{
		IdUsuario:     5,
		IdFormulario:  int(idFormulario),
		MontoAvaluado: montoAvaluado,
		FechaAvaluo:   time.Now(),
		Observacion:   new(string),
	}
	*avaluo.Observacion = "Avalúo generado automáticamente por el sistema"

	if err := s.repo.Create(avaluo); err != nil {
		return nil, err
	}

	return avaluo, nil
}

func (s *avaluoService) GetByID(id uint) (*Avaluo, error) {
	return s.repo.GetByID(id)
}

func (s *avaluoService) GetAll() ([]*Avaluo, error) {
	return s.repo.GetAll()
}

func (s *avaluoService) Delete(id uint) (string, error) {
	return s.repo.Delete(id)
}
