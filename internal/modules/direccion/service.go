package direccion

type DireccionService interface {
	Create(direccion *Direccion) error
	GetByID(id uint) (*Direccion, error)
	GetAll() ([]*Direccion, error)
	Update(direccion *Direccion) (string, error)
	Delete(id uint) (string, error)
}

type direccionService struct {
	repo DireccionRepository
}

func NewDireccionService(repo DireccionRepository) DireccionService {
	return &direccionService{repo}
}

func (s *direccionService) Create(direccion *Direccion) error {
	return s.repo.Create(direccion)
}

func (s *direccionService) GetByID(id uint) (*Direccion, error) {
	return s.repo.GetByID(id)
}

func (s *direccionService) GetAll() ([]*Direccion, error) {
	return s.repo.GetAll()
}

func (s *direccionService) Update(direccion *Direccion) (string, error) {
	return s.repo.Update(direccion)
}

func (s *direccionService) Delete(id uint) (string, error) {
	return s.repo.Delete(id)
}
