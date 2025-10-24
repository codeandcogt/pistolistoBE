package seccion

type SeccionService interface {
	Create(seccion *Seccion) error
	GetByID(id uint) (*Seccion, error)
	GetAll() ([]*Seccion, error)
	UpdateSeccion(id uint, updated *Seccion) (*Seccion, error)
	DeleteSeccion(id uint) (string, error)
}

type seccionService struct {
	repo SeccionRepository
}

func NewSeccionService(repo SeccionRepository) SeccionService {
	return &seccionService{repo}
}

func (s *seccionService) Create(seccion *Seccion) error {
	return s.repo.Create(seccion)
}

func (s *seccionService) GetByID(id uint) (*Seccion, error) {
	return s.repo.GetByID(id)
}

func (s *seccionService) GetAll() ([]*Seccion, error) {
	return s.repo.GetAll()
}

func (s *seccionService) UpdateSeccion(id uint, updated *Seccion) (*Seccion, error) {
	return s.repo.UpdateSeccion(id, updated)
}

func (s *seccionService) DeleteSeccion(id uint) (string, error) {
	return s.repo.DeleteSeccion(id)
}
