package almacen

type AlmacenService interface {
	Create(almacen *Almacen) error
	GetByID(id uint) (*Almacen, error)
	GetAllBySucursal(sucursalId uint) ([]*Almacen, error)
	UpdateAlmacen(id uint, updated *Almacen) (*Almacen, error)
	DeleteAlmacen(id uint) (string, error)
}

type almacenService struct {
	repo AlmacenRepository
}

func NewAlmacenService(repo AlmacenRepository) AlmacenService {
	return &almacenService{repo}
}

func (s *almacenService) Create(almacen *Almacen) error {
	return s.repo.Create(almacen)
}

func (s *almacenService) GetByID(id uint) (*Almacen, error) {
	return s.repo.GetByID(id)
}

func (s *almacenService) GetAllBySucursal(IdSucursal uint) ([]*Almacen, error) {
	return s.repo.GetAllBySucursal(IdSucursal)
}

func (s *almacenService) UpdateAlmacen(id uint, updated *Almacen) (*Almacen, error) {
	return s.repo.UpdateAlmacen(id, updated)
}

func (s *almacenService) DeleteAlmacen(id uint) (string, error) {
	return s.repo.DeleteAlmacen(id)
}
