package resenaEmpresa

type ResenaEmpresaService interface {
	CreateResena(resena *ResenaEmpresa) error
	GetResenaByID(id uint) (*ResenaEmpresa, error)
	GetAll() ([]*ResenaEmpresa, error)
	GetResenasByClienteID(clienteId uint) ([]*ResenaEmpresa, error)
	GetResenaByFacturaID(facturaId uint) (*ResenaEmpresa, error)
	UpdateResena(id uint, resena *ResenaEmpresa) error
	DeleteResena(id uint) error
	GetPromedioCalificacion() (float64, error)
}

type resenaEmpresaService struct {
	repo ResenaEmpresaRepository
}

func NewResenaEmpresaService(repo ResenaEmpresaRepository) ResenaEmpresaService {
	return &resenaEmpresaService{repo}
}

func (s *resenaEmpresaService) CreateResena(resena *ResenaEmpresa) error {
	return s.repo.Create(resena)
}

func (s *resenaEmpresaService) GetResenaByID(id uint) (*ResenaEmpresa, error) {
	return s.repo.GetByID(id)
}

func (s *resenaEmpresaService) GetAll() ([]*ResenaEmpresa, error) {
	return s.repo.GetAll()
}

func (s *resenaEmpresaService) GetResenasByClienteID(clienteId uint) ([]*ResenaEmpresa, error) {
	return s.repo.GetByClienteID(clienteId)
}

func (s *resenaEmpresaService) GetResenaByFacturaID(facturaId uint) (*ResenaEmpresa, error) {
	return s.repo.GetByFacturaID(facturaId)
}

func (s *resenaEmpresaService) UpdateResena(id uint, resena *ResenaEmpresa) error {
	return s.repo.Update(id, resena)
}

func (s *resenaEmpresaService) DeleteResena(id uint) error {
	return s.repo.Delete(id)
}

func (s *resenaEmpresaService) GetPromedioCalificacion() (float64, error) {
	return s.repo.GetPromedioCalificacion()
}
