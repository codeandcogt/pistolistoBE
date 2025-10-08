package permiso

type PermisoService interface {
	Create(permiso *Permiso) error
	GetAll() ([]*Permiso, error)
	GetByID(id uint) (*Permiso, error)
	Delete(id uint) (string, error)
	Update(permiso *Permiso) (string, error)
}

type permisoService struct {
	repo PermisoRepository
}

func NewPermisoService(repo PermisoRepository) PermisoService {
	return &permisoService{repo}
}

func (s *permisoService) Create(permiso *Permiso) error {
	return s.repo.Create(permiso)
}

func (s *permisoService) GetAll() ([]*Permiso, error) {
	return s.repo.GetAll()
}

func (s *permisoService) GetByID(id uint) (*Permiso, error) {
	return s.repo.GetByID(id)
}

func (s *permisoService) Delete(id uint) (string, error) {
	return s.repo.Delete(id)
}

func (s *permisoService) Update(permiso *Permiso) (string, error) {
	return s.repo.Update(permiso)
}
