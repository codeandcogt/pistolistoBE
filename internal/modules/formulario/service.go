package formulario

type FormularioService interface {
	CreateFormulario(formulario *Formulario) error
	GetFormularioByID(id uint) (*Formulario, error)
	GetAll() ([]*Formulario, error)
	UpdateFormulario(id uint, formulario *Formulario) error
	DeleteFormulario(id uint) error
}

type formularioService struct {
	repo FormularioRepository
}

func NewFormularioService(repo FormularioRepository) FormularioService {
	return &formularioService{repo}
}

func (s *formularioService) CreateFormulario(formulario *Formulario) error {
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
