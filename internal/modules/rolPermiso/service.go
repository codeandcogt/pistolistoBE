package rolpermiso

type RolPermisoService interface {
	Create(rolPermiso *RolPermiso) error
	GetAll() ([]*RolPermiso, error)
	GetByID(id uint) (*RolPermiso, error)
	Update(rolPermiso *RolPermiso) (string, error)
	Delete(id uint) (string, error)
}

type rolPermisoService struct {
	repo RolPermisoRepository
}

func NewRolPermosoService(repo RolPermisoRepository) RolPermisoService {
	return &rolPermisoService{repo}
}

func (s *rolPermisoService) Create(rolPermiso *RolPermiso) error {
	return s.repo.Create(rolPermiso)
}

func (s *rolPermisoService) GetAll() ([]*RolPermiso, error) {
	return s.repo.GetAll()
}

func (s *rolPermisoService) GetByID(id uint) (*RolPermiso, error) {
	return s.repo.GetByID(id)
}

func (s *rolPermisoService) Delete(id uint) (string, error) {
	return s.repo.Delete(id)
}

func (s *rolPermisoService) Update(rolPermiso *RolPermiso) (string, error) {
	return s.repo.Update(rolPermiso)
}
