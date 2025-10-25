package cuota

type CuotaService interface {
	Create(cuota *Cuota) error
	GetByID(id uint) (*Cuota, error)
	GetAll() ([]*Cuota, error)
	Update(id uint, cuota *Cuota) error
	Delete(id uint) error
}

type cuotaService struct {
	repo CuotaRepository
}

func NewCuotaService(repo CuotaRepository) CuotaService {
	return &cuotaService{repo}
}

func (s *cuotaService) Create(cuota *Cuota) error {
	return s.repo.Create(cuota)
}

func (s *cuotaService) GetByID(id uint) (*Cuota, error) {
	return s.repo.GetByID(id)
}

func (s *cuotaService) GetAll() ([]*Cuota, error) {
	return s.repo.GetAll()
}

func (s *cuotaService) Update(id uint, cuota *Cuota) error {
	return s.repo.Update(id, cuota)
}

func (s *cuotaService) Delete(id uint) error {
	return s.repo.Delete(id)
}
