package carrito

type CarritoService interface {
	CreateCarrito(carrito *Carrito) error
	GetCarritoByID(id uint) (*Carrito, error)
	GetCarritoByClienteID(clienteId uint) (*Carrito, error)
	GetAll() ([]*Carrito, error)
	UpdateCarrito(id uint, carrito *Carrito) error
	DeleteCarrito(id uint) error

	// CarritoItem methods
	AddItemToCarrito(item *CarritoItem) error
	UpdateCarritoItem(id uint, item *CarritoItem) error
	RemoveItemFromCarrito(id uint) error
	GetCarritoItems(carritoId uint) ([]*CarritoItem, error)
	CalcularTotales(carritoId uint) error
}

type carritoService struct {
	repo CarritoRepository
}

func NewCarritoService(repo CarritoRepository) CarritoService {
	return &carritoService{repo}
}

func (s *carritoService) CreateCarrito(carrito *Carrito) error {
	return s.repo.Create(carrito)
}

func (s *carritoService) GetCarritoByID(id uint) (*Carrito, error) {
	return s.repo.GetByID(id)
}

func (s *carritoService) GetCarritoByClienteID(clienteId uint) (*Carrito, error) {
	return s.repo.GetByClienteID(clienteId)
}

func (s *carritoService) GetAll() ([]*Carrito, error) {
	return s.repo.GetAll()
}

func (s *carritoService) UpdateCarrito(id uint, carrito *Carrito) error {
	return s.repo.Update(id, carrito)
}

func (s *carritoService) DeleteCarrito(id uint) error {
	return s.repo.Delete(id)
}

func (s *carritoService) AddItemToCarrito(item *CarritoItem) error {
	return s.repo.AddItem(item)
}

func (s *carritoService) UpdateCarritoItem(id uint, item *CarritoItem) error {
	return s.repo.UpdateItem(id, item)
}

func (s *carritoService) RemoveItemFromCarrito(id uint) error {
	return s.repo.RemoveItem(id)
}

func (s *carritoService) GetCarritoItems(carritoId uint) ([]*CarritoItem, error) {
	return s.repo.GetItemsByCarritoID(carritoId)
}

func (s *carritoService) CalcularTotales(carritoId uint) error {
	items, err := s.repo.GetItemsByCarritoID(carritoId)
	if err != nil {
		return err
	}

	var subtotal float64 = 0
	for _, item := range items {
		if item.Subtotal != nil {
			subtotal += *item.Subtotal
		}
	}

	carrito, err := s.repo.GetByID(carritoId)
	if err != nil {
		return err
	}

	descuento := float64(0)
	if carrito.Descuento != nil {
		descuento = *carrito.Descuento
	}

	impuesto := (subtotal - descuento) * 0.12 // IVA 12%
	total := subtotal - descuento + impuesto

	carrito.SubTotal = &subtotal
	carrito.Impuesto = &impuesto
	carrito.Total = &total

	return s.repo.Update(carritoId, carrito)
}
