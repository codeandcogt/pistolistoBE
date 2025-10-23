package piloto

type PilotoService interface {
	CreatePiloto(piloto *Piloto) error
	GetPilotoByID(id uint) (*Piloto, error)
	GetAll() ([]*Piloto, error)
	GetPilotosByAdministrativoID(administrativoId uint) ([]*Piloto, error)
	UpdatePiloto(id uint, piloto *Piloto) error
	DeletePiloto(id uint) error
}

type pilotoService struct {
	repo PilotoRepository
}

func NewPilotoService(repo PilotoRepository) PilotoService {
	return &pilotoService{repo}
}

func (s *pilotoService) CreatePiloto(piloto *Piloto) error {
	return s.repo.Create(piloto)
}

func (s *pilotoService) GetPilotoByID(id uint) (*Piloto, error) {
	return s.repo.GetByID(id)
}

func (s *pilotoService) GetAll() ([]*Piloto, error) {
	return s.repo.GetAll()
}

func (s *pilotoService) GetPilotosByAdministrativoID(administrativoId uint) ([]*Piloto, error) {
	return s.repo.GetByAdministrativoID(administrativoId)
}

func (s *pilotoService) UpdatePiloto(id uint, piloto *Piloto) error {
	return s.repo.Update(id, piloto)
}

func (s *pilotoService) DeletePiloto(id uint) error {
	return s.repo.Delete(id)
}
