package moneda

type MonedaService interface {
	CreateMoneda(moneda *Moneda) error
	GetMonedaByID(id uint) (*Moneda, error)
	GetAll() ([]*Moneda, error)
	UpdateMoneda(id uint, moneda *Moneda) error
	DeleteMoneda(id uint) error
	GetMonedaByCodigo(codigo string) (*Moneda, error)
}

type monedaService struct {
	repo MonedaRepository
}

func NewMonedaService(repo MonedaRepository) MonedaService {
	return &monedaService{repo}
}

func (s *monedaService) CreateMoneda(moneda *Moneda) error {
	return s.repo.Create(moneda)
}

func (s *monedaService) GetMonedaByID(id uint) (*Moneda, error) {
	return s.repo.GetByID(id)
}

func (s *monedaService) GetAll() ([]*Moneda, error) {
	return s.repo.GetAll()
}

func (s *monedaService) UpdateMoneda(id uint, moneda *Moneda) error {
	return s.repo.Update(id, moneda)
}

func (s *monedaService) DeleteMoneda(id uint) error {
	return s.repo.Delete(id)
}

func (s *monedaService) GetMonedaByCodigo(codigo string) (*Moneda, error) {
	return s.repo.GetByCodigo(codigo)
}
