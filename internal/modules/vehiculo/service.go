package vehiculo

type VehiculoService interface {
	CreateVehiculo(vehiculo *Vehiculo) error
	GetVehiculoByID(id uint) (*Vehiculo, error)
	GetAll() ([]*Vehiculo, error)
	GetVehiculosByPilotoID(pilotoId uint) ([]*Vehiculo, error)
	UpdateVehiculo(id uint, vehiculo *Vehiculo) error
	DeleteVehiculo(id uint) error
}

type vehiculoService struct {
	repo VehiculoRepository
}

func NewVehiculoService(repo VehiculoRepository) VehiculoService {
	return &vehiculoService{repo}
}

func (s *vehiculoService) CreateVehiculo(vehiculo *Vehiculo) error {
	return s.repo.Create(vehiculo)
}

func (s *vehiculoService) GetVehiculoByID(id uint) (*Vehiculo, error) {
	return s.repo.GetByID(id)
}

func (s *vehiculoService) GetAll() ([]*Vehiculo, error) {
	return s.repo.GetAll()
}

func (s *vehiculoService) GetVehiculosByPilotoID(pilotoId uint) ([]*Vehiculo, error) {
	return s.repo.GetByPilotoID(pilotoId)
}

func (s *vehiculoService) UpdateVehiculo(id uint, vehiculo *Vehiculo) error {
	return s.repo.Update(id, vehiculo)
}

func (s *vehiculoService) DeleteVehiculo(id uint) error {
	return s.repo.Delete(id)
}
