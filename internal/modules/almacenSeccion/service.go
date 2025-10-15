package almacenseccion

type AlmacenSeccionService interface {
	Create(seccion *AlmacenSeccion) error
	GetByID(id uint) (*AlmacenSeccion, error)
	GetByAlmacen(idAlmacen uint) ([]*AlmacenSeccion, error)
	UpdateAlmacenSeccion(id uint, updated *AlmacenSeccion) (*AlmacenSeccion, error)
	DeleteAlmacenSeccion(id uint) (string, error)
}

type almacenSeccionService struct {
	repo AlmacenSeccionRepository
}

func NewAlmacenSeccionService(repo AlmacenSeccionRepository) AlmacenSeccionService {
	return &almacenSeccionService{repo}
}

func (s *almacenSeccionService) Create(seccion *AlmacenSeccion) error {
	return s.repo.Create(seccion)
}

func (s *almacenSeccionService) GetByID(id uint) (*AlmacenSeccion, error) {
	return s.repo.GetByID(id)
}

func (s *almacenSeccionService) GetByAlmacen(idAlmacen uint) ([]*AlmacenSeccion, error) {
	return s.repo.GetByAlmacen(idAlmacen)
}

func (s *almacenSeccionService) UpdateAlmacenSeccion(id uint, updated *AlmacenSeccion) (*AlmacenSeccion, error) {
	return s.repo.UpdateAlmacenSeccion(id, updated)
}

func (s *almacenSeccionService) DeleteAlmacenSeccion(id uint) (string, error) {
	return s.repo.DeleteAlmacenSeccion(id)
}
