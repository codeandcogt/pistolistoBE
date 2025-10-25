package formulario

import (
	"fmt"
	"pistolistoBE/internal/modules/avaluo"
	"time"
)

type FormularioService interface {
	CreateFormulario(formulario *Formulario) error
	GetFormularioByID(id uint) (*Formulario, error)
	GetAll() ([]*Formulario, error)
	UpdateFormulario(id uint, formulario *Formulario) error
	DeleteFormulario(id uint) error

	// Agregado:
	GetFormularioLite(id uint) (*avaluo.FormularioLite, error)
}

type formularioService struct {
	repo FormularioRepository
}

func NewFormularioService(repo FormularioRepository) FormularioService {
	return &formularioService{repo}
}

func (s *formularioService) CreateFormulario(formulario *Formulario) error {
	formulario.NumeroFormulario = fmt.Sprintf("F-%d", time.Now().Unix())
	return s.repo.Create(formulario)
}

func (s *formularioService) GetFormularioByID(id uint) (*Formulario, error) {
	return s.repo.GetByID(id)
}

func (s *formularioService) GetAll() ([]*Formulario, error) {
	return s.repo.GetAll()
}

func (s *formularioService) UpdateFormulario(id uint, formulario *Formulario) error {
	return s.repo.Update(id, formulario)
}

func (s *formularioService) DeleteFormulario(id uint) error {
	return s.repo.Delete(id)
}

// --- Adaptador para Avaluo (cumple con la interfaz FormularioDataProvider) ---
func (s *formularioService) GetFormularioLite(id uint) (*avaluo.FormularioLite, error) {
	form, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &avaluo.FormularioLite{
		IdFormulario:    form.IdFormulario,
		IdArticulo:      form.IdArticulo,
		MontoSolicitado: form.MontoSolicitado,
		TipoOperacion:   form.TipoOperacion,
	}, nil
}
