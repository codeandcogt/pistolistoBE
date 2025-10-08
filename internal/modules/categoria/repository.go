package categoria

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type CategoriaRepository interface {
	Create(categoria *Categoria) error
	GetByID(id uint) (*Categoria, error)
	GetAll() ([]*Categoria, error)
	DeleteCategoria(id uint) (string, error)
}

type categoriaRepository struct {
	db *gorm.DB
}

func NewCategoria(db *gorm.DB) CategoriaRepository {
	return &categoriaRepository{db}
}

func (r *categoriaRepository) Create(categoria *Categoria) error {
	return r.db.Create(categoria).Error
}

func (r *categoriaRepository) GetByID(id uint) (*Categoria, error) {
	var categoria Categoria
	err := r.db.Where("estado = ?", true).First(&categoria, id).Error
	if err != nil {
		return nil, err
	}
	return &categoria, nil
}

func (r *categoriaRepository) GetAll() ([]*Categoria, error) {
	var categoria []*Categoria
	err := r.db.Where("estado = ?", true).Find(&categoria).Error

	if err != nil {
		return nil, err
	}
	return categoria, nil
}

func (r *categoriaRepository) DeleteCategoria(id uint) (string, error) {
	result := r.db.Model(&Categoria{}).Where("id_categoria = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}
