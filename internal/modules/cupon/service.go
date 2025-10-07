package cupon

type CuponService interface {
	CreateCupon(cupon *Cupon) error
	GetCuponByID(id uint) (*Cupon, error)
	GetAll() ([]*Cupon, error)
	UpdateCupon(id uint, cupon *Cupon) error
	DeleteCupon(id uint) error
	GetCuponByCodigo(codigo string) (*Cupon, error)
}

type cuponService struct {
	repo CuponRepository
}

func NewCuponService(repo CuponRepository) CuponService {
	return &cuponService{repo}
}

func (s *cuponService) CreateCupon(cupon *Cupon) error {
	return s.repo.Create(cupon)
}

func (s *cuponService) GetCuponByID(id uint) (*Cupon, error) {
	return s.repo.GetByID(id)
}

func (s *cuponService) GetAll() ([]*Cupon, error) {
	return s.repo.GetAll()
}

func (s *cuponService) UpdateCupon(id uint, cupon *Cupon) error {
	return s.repo.Update(id, cupon)
}

func (s *cuponService) DeleteCupon(id uint) error {
	return s.repo.Delete(id)
}

func (s *cuponService) GetCuponByCodigo(codigo string) (*Cupon, error) {
	return s.repo.GetByCodigo(codigo)
}
