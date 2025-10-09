package administrativo

type AdminService interface {
	Create(admin *Administrativo) error
	GetAll() ([]*Administrativo, error)
	GetByID(id uint) (*Administrativo, error)
	Delete(id uint) (string, error)
	Update(admin *Administrativo) (string, error)
	FirstAuth(id uint) (string, error)
}

type adminService struct {
	repo AdminRepository
}

func NewAdministrativoService(repo AdminRepository) AdminService {
	return &adminService{repo}
}

func (s *adminService) Create(permiso *Administrativo) error {
	return s.repo.Create(permiso)
}

func (s *adminService) GetAll() ([]*Administrativo, error) {
	return s.repo.GetAll()
}

func (s *adminService) GetByID(id uint) (*Administrativo, error) {
	return s.repo.GetByID(id)
}

func (s *adminService) Delete(id uint) (string, error) {
	return s.repo.Delete(id)
}

func (s *adminService) Update(permiso *Administrativo) (string, error) {
	return s.repo.Update(permiso)
}

func (s *adminService) FirstAuth(id uint) (string, error) {
	return s.repo.FirstAuth(id)
}
