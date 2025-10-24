package resenaEmpresa

import (
	"gorm.io/gorm"
)

type ResenaEmpresaRepository interface {
	Create(resena *ResenaEmpresa) error
	GetByID(id uint) (*ResenaEmpresa, error)
	GetAll() ([]*ResenaEmpresa, error)
	GetByClienteID(clienteId uint) ([]*ResenaEmpresa, error)
	GetByFacturaID(facturaId uint) (*ResenaEmpresa, error)
	Update(id uint, resena *ResenaEmpresa) error
	Delete(id uint) error
	GetPromedioCalificacion() (float64, error)
}

type resenaEmpresaRepository struct {
	db *gorm.DB
}

func NewResenaEmpresaRepository(db *gorm.DB) ResenaEmpresaRepository {
	return &resenaEmpresaRepository{db}
}

func (r *resenaEmpresaRepository) Create(resena *ResenaEmpresa) error {
	return r.db.Create(resena).Error
}

func (r *resenaEmpresaRepository) GetByID(id uint) (*ResenaEmpresa, error) {
	var resena ResenaEmpresa
	err := r.db.Where("id_resena_empresa = ? AND estado = ?", id, true).First(&resena).Error
	if err != nil {
		return nil, err
	}
	return &resena, nil
}

func (r *resenaEmpresaRepository) GetAll() ([]*ResenaEmpresa, error) {
	var resenas []*ResenaEmpresa
	err := r.db.Where("estado = ?", true).Order("fecha_creacion DESC").Find(&resenas).Error
	if err != nil {
		return nil, err
	}
	return resenas, nil
}

func (r *resenaEmpresaRepository) GetByClienteID(clienteId uint) ([]*ResenaEmpresa, error) {
	var resenas []*ResenaEmpresa
	err := r.db.Where("cliente_id = ? AND estado = ?", clienteId, true).Order("fecha_creacion DESC").Find(&resenas).Error
	if err != nil {
		return nil, err
	}
	return resenas, nil
}

func (r *resenaEmpresaRepository) GetByFacturaID(facturaId uint) (*ResenaEmpresa, error) {
	var resena ResenaEmpresa
	err := r.db.Where("factura_id = ? AND estado = ?", facturaId, true).First(&resena).Error
	if err != nil {
		return nil, err
	}
	return &resena, nil
}

func (r *resenaEmpresaRepository) Update(id uint, resena *ResenaEmpresa) error {
	return r.db.Model(&ResenaEmpresa{}).Where("id_resena_empresa = ? AND estado = ?", id, true).Updates(resena).Error
}

func (r *resenaEmpresaRepository) Delete(id uint) error {
	return r.db.Model(&ResenaEmpresa{}).Where("id_resena_empresa = ? AND estado = ?", id, true).Update("estado", false).Error
}

func (r *resenaEmpresaRepository) GetPromedioCalificacion() (float64, error) {
	var promedio float64
	err := r.db.Model(&ResenaEmpresa{}).Where("estado = ?", true).Select("AVG(calificacion)").Scan(&promedio).Error
	return promedio, err
}
