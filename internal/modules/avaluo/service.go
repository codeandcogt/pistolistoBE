package avaluo

import (
	"errors"
	"fmt"
	"math/rand"
	"pistolistoBE/internal/modules/articulo"
	"pistolistoBE/internal/modules/contrato"
	"time"
)

// FormularioLite es una vista ligera del formulario para evitar dependencias cíclicas
type FormularioLite struct {
	IdFormulario    uint
	IdArticulo      int
	MontoSolicitado float64
	TipoOperacion   string
}

// FormularioDataProvider define cómo Avaluo obtiene los datos del formulario
type FormularioDataProvider interface {
	GetFormularioLite(id uint) (*FormularioLite, error)
}

type AvaluoService interface {
	CrearAvaluoManual(avaluo *Avaluo) error
	CrearAvaluoAutomatico(idFormulario uint) (*Avaluo, error)
	GetByID(id uint) (*Avaluo, error)
	GetAll() ([]*Avaluo, error)
	Delete(id uint) (string, error)

	GetContratoByAvaluo(idAvaluo uint) (*contrato.Contrato, error)
	GetArticuloByFormulario(idFormulario uint) (*articulo.Articulo, error)
}

type avaluoService struct {
	repo         AvaluoRepository
	formProvider FormularioDataProvider
	contratoSrv  contrato.ContratoService
	articuloSrv  articulo.ArticuloService
}

// Constructor limpio
func NewAvaluoService(
	repo AvaluoRepository,
	formProvider FormularioDataProvider,
	contratoSrv contrato.ContratoService,
	articuloSrv articulo.ArticuloService,
) AvaluoService {
	return &avaluoService{
		repo:         repo,
		formProvider: formProvider,
		contratoSrv:  contratoSrv,
		articuloSrv:  articuloSrv,
	}
}

// Crear avalúo manual
func (s *avaluoService) CrearAvaluoManual(avaluo *Avaluo) error {
	avaluo.FechaAvaluo = time.Now()
	return s.repo.Create(avaluo)
}

// Crear avalúo automático según datos del formulario
func (s *avaluoService) CrearAvaluoAutomatico(idFormulario uint) (*Avaluo, error) {
	form, err := s.formProvider.GetFormularioLite(idFormulario)
	if err != nil {
		return nil, errors.New("no se encontró el formulario")
	}

	// Calcular monto según tipo de operación
	factor := 1.0
	switch form.TipoOperacion {
	case "compra":
		factor = 1.05
	case "venta":
		factor = 0.95
	case "empenio":
		factor = 1.02
	}

	// Variación aleatoria
	if rand.Float64() > 0.7 {
		factor += 0.03
	}

	montoAvaluado := form.MontoSolicitado * factor

	// Crear avalúo
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

	// Crear contrato automáticamente
	_, err = s.contratoSrv.CrearContratoAutomatico(avaluo.IdAvaluo, montoAvaluado)
	if err != nil {
		fmt.Println("⚠️ Error al crear contrato automático:", err)
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

// Obtener contrato relacionado al avalúo
func (s *avaluoService) GetContratoByAvaluo(idAvaluo uint) (*contrato.Contrato, error) {
	return s.contratoSrv.GetByAvaluo(idAvaluo)
}

// Obtener artículo vinculado al formulario
func (s *avaluoService) GetArticuloByFormulario(idFormulario uint) (*articulo.Articulo, error) {
	form, err := s.formProvider.GetFormularioLite(idFormulario)
	if err != nil {
		return nil, err
	}

	if form.IdArticulo == 0 {
		return nil, fmt.Errorf("el formulario no tiene un artículo asociado")
	}

	articuloData, err := s.articuloSrv.GetByID(uint(form.IdArticulo))
	if err != nil {
		return nil, fmt.Errorf("no se encontró el artículo asociado: %v", err)
	}

	return articuloData, nil
}
