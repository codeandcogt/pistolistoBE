package producto

type ProductoService interface {
	Create(producto *Producto) error
	GetByID(id uint) (*Producto, error)
	GetAll() ([]*Producto, error)
	Update(producto *Producto) (string, error)
	Delete(id uint) (string, error)
}

type productoService struct {
	repo ProductoRepository
}

func NewProductoService(repo ProductoRepository) ProductoService {
	return &productoService{repo}
}

func (s *productoService) Create(producto *Producto) error {
	return s.repo.Create(producto)
}

func (s *productoService) GetByID(id uint) (*Producto, error) {
	return s.repo.GetByID(id)
}

func (s *productoService) GetAll() ([]*Producto, error) {
	return s.repo.GetAll()
}

func (s *productoService) Update(producto *Producto) (string, error) {
	return s.repo.Update(producto)
}

func (s *productoService) Delete(id uint) (string, error) {
	return s.repo.Delete(id)
}
