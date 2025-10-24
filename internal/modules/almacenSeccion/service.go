package almacenseccion

type AlmacenSeccionService interface {
	Create(almseccion *AlmacenSeccion) error
	GetByID(id uint) (*AlmacenSeccion, error)
	GetByAlmacen(IdAlmacen uint) ([]*AlmacenSeccion, error)
	UpdateAlmacenSeccion(id uint, updated *AlmacenSeccion) (*AlmacenSeccion, error)
	DeleteAlmacenSeccion(id uint) (string, error)
}

type almacenSeccionService struct {
	repo AlmacenSeccionRepository
}

func NewAlmacenSeccionService(repo AlmacenSeccionRepository) AlmacenSeccionService {
	return &almacenSeccionService{repo}
}

func (s *almacenSeccionService) Create(almseccion *AlmacenSeccion) error {
	return s.repo.Create(almseccion)
}

func (s *almacenSeccionService) GetByID(id uint) (*AlmacenSeccion, error) {
	return s.repo.GetByID(id)
}

func (s *almacenSeccionService) GetByAlmacen(IdAlmacen uint) ([]*AlmacenSeccion, error) {
	return s.repo.GetByAlmacen(IdAlmacen)
}

func (s *almacenSeccionService) UpdateAlmacenSeccion(id uint, updated *AlmacenSeccion) (*AlmacenSeccion, error) {
	return s.repo.UpdateAlmacenSeccion(id, updated)
}

func (s *almacenSeccionService) DeleteAlmacenSeccion(id uint) (string, error) {
	return s.repo.DeleteAlmacenSeccion(id)
}
