package rol

type RolService interface {
	CreateRol(rol *Rol) error
	GetAll() ([]*Rol, error)
}

type rolService struct {
	repo RolRepository
}

func NewRolService(repo RolRepository) RolService {
	return &rolService{repo}
}

func (s *rolService) CreateRol(rol *Rol) error {
	return s.repo.Create(rol)
}

func (s *rolService) GetAll() ([]*Rol, error) {
	return s.repo.GetAll()
}
