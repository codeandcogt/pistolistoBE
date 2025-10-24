package ruta

type RutaService interface {
	Create(ruta *Ruta) error
	GetByID(id uint) (*Ruta, error)
	GetAll() ([]*Ruta, error)
	Update(ruta *Ruta) (string, error)
	Delete(id uint) (string, error)
	CambiarEstado(idRuta uint, idEstadoRuta int) error
}

type rutaService struct {
	repo RutaRepository
}

func NewRutaService(repo RutaRepository) RutaService {
	return &rutaService{repo}
}

func (s *rutaService) Create(ruta *Ruta) error {
	return s.repo.Create(ruta)
}

func (s *rutaService) GetByID(id uint) (*Ruta, error) {
	return s.repo.GetByID(id)
}

func (s *rutaService) GetAll() ([]*Ruta, error) {
	return s.repo.GetAll()
}

func (s *rutaService) Update(ruta *Ruta) (string, error) {
	return s.repo.Update(ruta)
}

func (s *rutaService) Delete(id uint) (string, error) {
	return s.repo.Delete(id)
}

func (s *rutaService) CambiarEstado(idRuta uint, idEstadoRuta int) error {
	return s.repo.UpdateEstado(idRuta, idEstadoRuta)
}
