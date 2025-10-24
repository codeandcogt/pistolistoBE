package logUbicacion

import (
	"pistolistoBE/internal/modules/logUbicacionTiempoReal"
)

type LogUbicacionService interface {
	RegistrarUbicacion(log *LogUbicacion) error
	GetByPiloto(pilotoId uint) ([]*LogUbicacion, error)
}

type logUbicacionService struct {
	repo           LogUbicacionRepository
	tiempoRealRepo logUbicacionTiempoReal.LogUbicacionTiempoRealRepository
}

func NewLogUbicacionService(repo LogUbicacionRepository, tiempoRealRepo logUbicacionTiempoReal.LogUbicacionTiempoRealRepository) LogUbicacionService {
	return &logUbicacionService{repo, tiempoRealRepo}
}

func (s *logUbicacionService) RegistrarUbicacion(log *LogUbicacion) error {
	// 1️⃣ Guardar el histórico
	if err := s.repo.Create(log); err != nil {
		return err
	}

	// 2️⃣ Actualizar ubicación actual del piloto
	if err := s.tiempoRealRepo.Upsert(log.PilotoId, log.Latitud, log.Longitud); err != nil {
		return err
	}

	// 3️⃣ Emitir actualización a los clientes WebSocket
	logUbicacionTiempoReal.BroadcastUbicacion(log.PilotoId, log.Latitud, log.Longitud)

	return nil
}

func (s *logUbicacionService) GetByPiloto(pilotoId uint) ([]*LogUbicacion, error) {
	return s.repo.GetByPiloto(pilotoId)
}
