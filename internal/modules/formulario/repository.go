package formulario

import (
	"gorm.io/gorm"
)

type FormularioRepository interface {
	Create(formulario *Formulario) error
	GetByID(id uint) (*Formulario, error)
	GetAll() ([]*Formulario, error)
	Update(id uint, formulario *Formulario) error
	Delete(id uint) error
}

type formularioRepository struct {
	db *gorm.DB
}

func NewFormularioRepository(db *gorm.DB) FormularioRepository {
	return &formularioRepository{db}
}

func (r *formularioRepository) Create(formulario *Formulario) error {
	return r.db.Create(formulario).Error
}

func (r *formularioRepository) GetByID(id uint) (*Formulario, error) {
	var formulario Formulario
	err := r.db.Where("id_formulario = ? AND estado = ?", id, true).First(&formulario).Error
	if err != nil {
		return nil, err
	}
	return &formulario, nil
}

func (r *formularioRepository) GetAll() ([]*Formulario, error) {
	var formularios []*Formulario
	err := r.db.Where("estado = ?", true).Order("fecha_creacion DESC").Find(&formularios).Error
	if err != nil {
		return nil, err
	}
	return formularios, nil
}

func (r *formularioRepository) Update(id uint, formulario *Formulario) error {
	return r.db.Model(&Formulario{}).Where("id_formulario = ? AND estado = ?", id, true).Updates(formulario).Error
}

func (r *formularioRepository) Delete(id uint) error {
	return r.db.Model(&Formulario{}).Where("id_formulario = ? AND estado = ?", id, true).Update("estado", false).Error
}
