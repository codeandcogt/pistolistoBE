package municipality

type MunicipalityService interface {
	Create(municipality *Municipality) error
	GetByID(id uint) (*Municipality, error)
	GetAll() ([]*Municipality, error)
	Update(municipality *Municipality) (string, error)
	Delete(id uint) (string, error)
}

type municipalityService struct {
	repo MunicipalityRepository
}

func NewMunicipalityService(repo MunicipalityRepository) MunicipalityService {
	return &municipalityService{repo}
}

func (s *municipalityService) Create(municipality *Municipality) error {
	return s.repo.Create(municipality)
}

func (s *municipalityService) GetByID(id uint) (*Municipality, error) {
	return s.repo.GetByID(id)
}

func (s *municipalityService) GetAll() ([]*Municipality, error) {
	return s.repo.GetAll()
}

func (s *municipalityService) Update(municipality *Municipality) (string, error) {
	return s.repo.Update(municipality)
}

func (s *municipalityService) Delete(id uint) (string, error) {
	return s.repo.Delete(id)
}
