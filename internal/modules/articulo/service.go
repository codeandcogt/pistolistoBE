package articulo

type ArticuloService interface {
	Create(articulo *Articulo) error
	GetByID(id uint) (*Articulo, error)
	GetAll() ([]*Articulo, error)
	Update(articulo *Articulo) (string, error)
	Delete(id uint) (string, error)
}

type articuloService struct {
	repo ArticuloRepository
}

func NewArticuloService(repo ArticuloRepository) ArticuloService {
	return &articuloService{repo}
}

func (s *articuloService) Create(articulo *Articulo) error {
	return s.repo.Create(articulo)
}

func (s *articuloService) GetByID(id uint) (*Articulo, error) {
	return s.repo.GetByID(id)
}

func (s *articuloService) GetAll() ([]*Articulo, error) {
	return s.repo.GetAll()
}

func (s *articuloService) Update(articulo *Articulo) (string, error) {
	return s.repo.Update(articulo)
}

func (s *articuloService) Delete(id uint) (string, error) {
	return s.repo.Delete(id)
}
