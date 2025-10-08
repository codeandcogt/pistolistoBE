package rol

type RolService interface {
	CreateRol(rol *Rol) error
	GetAll() ([]*Rol, error)
	GetByID(id uint) (*Rol, error)
	Delete(id uint) (string, error)
	Update(rol *Rol) (string, error)
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

func (s *rolService) GetByID(id uint) (*Rol, error) {
	return s.repo.GetByID(id)
}

func (s *rolService) Delete(id uint) (string, error) {
	return s.repo.Delete(id)
}

func (s *rolService) Update(rol *Rol) (string, error) {
	return s.repo.Update(rol)
}
