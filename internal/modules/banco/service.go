package banco

type BancoService interface {
	CreateBanco(banco *Banco) error
	GetBancoByID(id uint) (*Banco, error)
	GetAll() ([]*Banco, error)
	UpdateBanco(id uint, banco *Banco) error
	DeleteBanco(id uint) error
	GetBancoByCodigo(codigo string) (*Banco, error)
	GetBancosByTipo(tipo string) ([]*Banco, error)
}

type bancoService struct {
	repo BancoRepository
}

func NewBancoService(repo BancoRepository) BancoService {
	return &bancoService{repo}
}

func (s *bancoService) CreateBanco(banco *Banco) error {
	return s.repo.Create(banco)
}

func (s *bancoService) GetBancoByID(id uint) (*Banco, error) {
	return s.repo.GetByID(id)
}

func (s *bancoService) GetAll() ([]*Banco, error) {
	return s.repo.GetAll()
}

func (s *bancoService) UpdateBanco(id uint, banco *Banco) error {
	return s.repo.Update(id, banco)
}

func (s *bancoService) DeleteBanco(id uint) error {
	return s.repo.Delete(id)
}

func (s *bancoService) GetBancoByCodigo(codigo string) (*Banco, error) {
	return s.repo.GetByCodigo(codigo)
}

func (s *bancoService) GetBancosByTipo(tipo string) ([]*Banco, error) {
	return s.repo.GetByTipo(tipo)
}
