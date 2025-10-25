package prestamo

type PrestamoService interface {
	CreatePrestamo(prestamo *Prestamo) error
	GetPrestamoByID(id uint) (*Prestamo, error)
	GetAll() ([]*Prestamo, error)
	GetPrestamosByClienteID(clienteId uint) ([]*Prestamo, error)
	UpdatePrestamo(id uint, prestamo *Prestamo) error
	DeletePrestamo(id uint) error
}

type prestamoService struct {
	repo PrestamoRepository
}

func NewPrestamoService(repo PrestamoRepository) PrestamoService {
	return &prestamoService{repo}
}

func (s *prestamoService) CreatePrestamo(prestamo *Prestamo) error {
	return s.repo.Create(prestamo)
}

func (s *prestamoService) GetPrestamoByID(id uint) (*Prestamo, error) {
	return s.repo.GetByID(id)
}

func (s *prestamoService) GetAll() ([]*Prestamo, error) {
	return s.repo.GetAll()
}

func (s *prestamoService) GetPrestamosByClienteID(clienteId uint) ([]*Prestamo, error) {
	return s.repo.GetByClienteID(clienteId)
}

func (s *prestamoService) UpdatePrestamo(id uint, prestamo *Prestamo) error {
	return s.repo.Update(id, prestamo)
}

func (s *prestamoService) DeletePrestamo(id uint) error {
	return s.repo.Delete(id)
}
