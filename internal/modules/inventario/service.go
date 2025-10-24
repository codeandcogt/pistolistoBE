package inventario

type InventarioService interface {
	CreateInventario(inventario *Inventario) error
	GetInventarioByID(id uint) (*Inventario, error)
	GetAllByAlmacen(IdAlmacen uint) ([]*Inventario, error)
	GetByTipo(IdAlmacen uint, tipoItem string) ([]*Inventario, error)
	GetByEstado(IdAlmacen uint, estadoInventario string) ([]*Inventario, error)
	GetProductosDisponibles(IdAlmacen uint) ([]*Inventario, error)
	GetArticulosEmpenados(IdAlmacen uint) ([]*Inventario, error)
	UpdateInventario(id uint, updated *Inventario) (*Inventario, error)
	DeleteInventario(id uint) (string, error)
}

type inventarioService struct {
	repo InventarioRepository
}

func NewInventarioService(repo InventarioRepository) InventarioService {
	return &inventarioService{repo}
}

func (s *inventarioService) CreateInventario(inventario *Inventario) error {
	return s.repo.Create(inventario)
}

func (s *inventarioService) GetInventarioByID(id uint) (*Inventario, error) {
	return s.repo.GetByID(id)
}

func (s *inventarioService) GetAllByAlmacen(IdAlmacen uint) ([]*Inventario, error) {
	return s.repo.GetAllByAlmacen(IdAlmacen)
}

func (s *inventarioService) GetByTipo(IdAlmacen uint, tipoItem string) ([]*Inventario, error) {
	return s.repo.GetByTipo(IdAlmacen, tipoItem)
}

func (s *inventarioService) GetByEstado(IdAlmacen uint, estadoInventario string) ([]*Inventario, error) {
	return s.repo.GetByEstado(IdAlmacen, estadoInventario)
}

func (s *inventarioService) GetProductosDisponibles(IdAlmacen uint) ([]*Inventario, error) {
	return s.repo.GetProductosDisponibles(IdAlmacen)
}

func (s *inventarioService) GetArticulosEmpenados(IdAlmacen uint) ([]*Inventario, error) {
	return s.repo.GetArticulosEmpenados(IdAlmacen)
}

func (s *inventarioService) UpdateInventario(id uint, updated *Inventario) (*Inventario, error) {
	return s.repo.UpdateInventario(id, updated)
}

func (s *inventarioService) DeleteInventario(id uint) (string, error) {
	return s.repo.DeleteInventario(id)
}
