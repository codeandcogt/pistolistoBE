package estadoRuta

type EstadoRutaService interface {
	Create(estado *EstadoRuta) error
	GetByID(id uint) (*EstadoRuta, error)
	GetAll() ([]*EstadoRuta, error)
	Update(estado *EstadoRuta) (string, error)
	Delete(id uint) (string, error)
}

type estadoRutaService struct {
	repo EstadoRutaRepository
}

func NewEstadoRutaService(repo EstadoRutaRepository) EstadoRutaService {
	return &estadoRutaService{repo}
}

func (s *estadoRutaService) Create(estado *EstadoRuta) error {
	return s.repo.Create(estado)
}

func (s *estadoRutaService) GetByID(id uint) (*EstadoRuta, error) {
	return s.repo.GetByID(id)
}

func (s *estadoRutaService) GetAll() ([]*EstadoRuta, error) {
	return s.repo.GetAll()
}

func (s *estadoRutaService) Update(estado *EstadoRuta) (string, error) {
	return s.repo.Update(estado)
}

func (s *estadoRutaService) Delete(id uint) (string, error) {
	return s.repo.Delete(id)
}
