package articulo

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type ArticuloRepository interface {
	Create(articulo *Articulo) error
	GetByID(id uint) (*Articulo, error)
	GetAll() ([]*Articulo, error)
	Update(articulo *Articulo) (string, error)
	Delete(id uint) (string, error)
}

type articuloRepository struct {
	db *gorm.DB
}

func NewArticuloRepository(db *gorm.DB) ArticuloRepository {
	return &articuloRepository{db}
}

func (r *articuloRepository) Create(articulo *Articulo) error {
	return r.db.Create(articulo).Error
}

func (r *articuloRepository) GetByID(id uint) (*Articulo, error) {
	var articulo Articulo
	err := r.db.Where("estado = ?", true).First(&articulo, id).Error
	if err != nil {
		return nil, err
	}
	return &articulo, nil
}

func (r *articuloRepository) GetAll() ([]*Articulo, error) {
	var articulos []*Articulo
	err := r.db.Where("estado = ?", true).Find(&articulos).Error
	if err != nil {
		return nil, err
	}
	return articulos, nil
}

func (r *articuloRepository) Update(articulo *Articulo) (string, error) {
	result := r.db.Model(&Articulo{}).
		Where("id_articulo = ? AND estado = ?", articulo.IdArticulo, true).
		Updates(articulo)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_UPDATED, nil
}

func (r *articuloRepository) Delete(id uint) (string, error) {
	result := r.db.Model(&Articulo{}).Where("id_articulo = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}
