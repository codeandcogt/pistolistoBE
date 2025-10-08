package subCategory

type SubCategoryService interface {
	Create(subCategory *SubCategory) error
	GetByID(id uint) (*SubCategory, error)
	GetAll() ([]*SubCategory, error)
	Update(subCategory *SubCategory) (string, error)
	Delete(id uint) (string, error)
}

type subCategoryService struct {
	repo SubCategoryRepository
}

func NewSubCategoryService(repo SubCategoryRepository) SubCategoryService {
	return &subCategoryService{repo}
}

func (s *subCategoryService) Create(subCategory *SubCategory) error {
	return s.repo.Create(subCategory)
}

func (s *subCategoryService) GetByID(id uint) (*SubCategory, error) {
	return s.repo.GetByID(id)
}

func (s *subCategoryService) GetAll() ([]*SubCategory, error) {
	return s.repo.GetAll()
}

func (s *subCategoryService) Update(subCategory *SubCategory) (string, error) {
	return s.repo.Update(subCategory)
}

func (s *subCategoryService) Delete(id uint) (string, error) {
	return s.repo.Delete(id)
}
