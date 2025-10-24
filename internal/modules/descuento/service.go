package descuento

type DescuentoService interface {
	CreateDescuento(descuento *Descuento) error
	GetDescuentoByID(id uint) (*Descuento, error)
	GetAll() ([]*Descuento, error)
	UpdateDescuento(id uint, updated *Descuento) (*Descuento, error)
	DeleteDescuento(id uint) (string, error)
}

type descuentoService struct {
	repo DescuentoRepository
}

func NewDescuentoService(repo DescuentoRepository) DescuentoService {
	return &descuentoService{repo}
}

func (s *descuentoService) CreateDescuento(descuento *Descuento) error {
	return s.repo.Create(descuento)
}

func (s *descuentoService) GetDescuentoByID(id uint) (*Descuento, error) {
	return s.repo.GetByID(id)
}

func (s *descuentoService) GetAll() ([]*Descuento, error) {
	return s.repo.GetAll()
}

func (s *descuentoService) UpdateDescuento(id uint, updated *Descuento) (*Descuento, error) {
	return s.repo.UpdateDescuento(id, updated)
}

func (s *descuentoService) DeleteDescuento(id uint) (string, error) {
	return s.repo.DeleteDescuento(id)
}
