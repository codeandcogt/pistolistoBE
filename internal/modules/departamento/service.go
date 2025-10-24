package departamento

type DepartamentoService interface {
	CreateDepartamento(departamento *Departamento) error
	GetDepartamentoByID(id uint) (*Departamento, error)
	GetAll() ([]*Departamento, error)
	UpdateDepartamento(id uint, updated *Departamento) (*Departamento, error)
	DeleteDepartamento(id uint) (string, error)
}

type departamentoService struct {
	repo DepartamentoRepository
}

func NewDepartamentoService(repo DepartamentoRepository) DepartamentoService {
	return &departamentoService{repo}
}

func (s *departamentoService) CreateDepartamento(departamento *Departamento) error {
	return s.repo.Create(departamento)
}

func (s *departamentoService) GetDepartamentoByID(id uint) (*Departamento, error) {
	return s.repo.GetByID(id)
}

func (s *departamentoService) GetAll() ([]*Departamento, error) {
	return s.repo.GetAll()
}

func (s *departamentoService) UpdateDepartamento(id uint, updated *Departamento) (*Departamento, error) {
	return s.repo.UpdateDepartamento(id, updated)
}

func (s *departamentoService) DeleteDepartamento(id uint) (string, error) {
	return s.repo.DeleteDepartamento(id)
}
