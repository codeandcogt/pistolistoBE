package cliente

type ClienteService interface {
	CreateCliente(client *Cliente) error
}

type clienteService struct {
	repo ClienteRepository
}

func NewClientService(repo ClienteRepository) ClienteService {
	return &clienteService{repo}
}

func (s *clienteService) CreateCliente(client *Cliente) error {
	return s.repo.Create(client)
}
