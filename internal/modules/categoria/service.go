package categoria

type CategoriaService interface {
	CreateCategoria(categoria *Categoria) error
	GetCategoriaByID(id uint) (*Categoria, error)
	GetAll() ([]*Categoria, error)
	DeleteCategoria(id uint) (string, error)
}

type categoriaService struct {
	repo CategoriaRepository
}

func NewCategoriaService(repo CategoriaRepository) CategoriaService {
	return &categoriaService{repo}
}

func (s *categoriaService) CreateCategoria(categoria *Categoria) error {
	return s.repo.Create(categoria)
}

func (s *categoriaService) GetCategoriaByID(id uint) (*Categoria, error) {
	return s.repo.GetByID(id)
}

func (s *categoriaService) GetAll() ([]*Categoria, error) {
	return s.repo.GetAll()
}

func (s *categoriaService) DeleteCategoria(id uint) (string, error) {
	return s.repo.DeleteCategoria(id)
}
