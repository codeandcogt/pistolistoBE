package estadoPedido

type EstadoPedidoService interface {
	Create(estado *EstadoPedido) error
	GetByID(id uint) (*EstadoPedido, error)
	GetAll() ([]*EstadoPedido, error)
	Update(estado *EstadoPedido) (string, error)
	Delete(id uint) (string, error)
}

type estadoPedidoService struct {
	repo EstadoPedidoRepository
}

func NewEstadoPedidoService(repo EstadoPedidoRepository) EstadoPedidoService {
	return &estadoPedidoService{repo}
}

func (s *estadoPedidoService) Create(estado *EstadoPedido) error {
	return s.repo.Create(estado)
}

func (s *estadoPedidoService) GetByID(id uint) (*EstadoPedido, error) {
	return s.repo.GetByID(id)
}

func (s *estadoPedidoService) GetAll() ([]*EstadoPedido, error) {
	return s.repo.GetAll()
}

func (s *estadoPedidoService) Update(estado *EstadoPedido) (string, error) {
	return s.repo.Update(estado)
}

func (s *estadoPedidoService) Delete(id uint) (string, error) {
	return s.repo.Delete(id)
}
