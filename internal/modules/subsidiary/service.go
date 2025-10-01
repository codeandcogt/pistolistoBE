package subsidiary

type SubsidiaryService interface {
	Create(subsidiary *Subsidiary) error
	GetByID(id uint) (*Subsidiary, error)
	GetAll() ([]*Subsidiary, error)
	Update(subsidiary *Subsidiary) error
	Delete(id uint) (string, error)
}

type subsidiaryService struct {
	repo SubsidiaryRepository
}

func NewSubsidiaryService(repo SubsidiaryRepository) SubsidiaryService {
	return &subsidiaryService{repo}
}

func (s *subsidiaryService) Create(subsidiary *Subsidiary) error {
	return s.repo.Create(subsidiary)
}

func (s *subsidiaryService) GetByID(id uint) (*Subsidiary, error) {
	return s.repo.GetByID(id)
}

func (s *subsidiaryService) GetAll() ([]*Subsidiary, error) {
	return s.repo.GetAll()
}

func (s *subsidiaryService) Update(subsidiary *Subsidiary) error {
	return s.repo.Update(subsidiary)
}

func (s *subsidiaryService) Delete(id uint) (string, error) {
	return s.repo.Delete(id)
}
