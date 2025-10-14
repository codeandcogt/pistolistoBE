package subCategory

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type SubCategoryRepository interface {
	Create(subCategory *SubCategory) error
	GetByID(id uint) (*SubCategory, error)
	GetAll() ([]*SubCategory, error)
	Update(subCategory *SubCategory) (string, error)
	Delete(id uint) (string, error)
}

type subCategoryRepository struct {
	db *gorm.DB
}

func NewSubCategoryRepository(db *gorm.DB) SubCategoryRepository {
	return &subCategoryRepository{db}
}

func (r *subCategoryRepository) Create(subCategory *SubCategory) error {
	return r.db.Create(subCategory).Error
}

func (r *subCategoryRepository) GetByID(id uint) (*SubCategory, error) {
	var subCategory SubCategory
	err := r.db.Where("estado = ?", true).First(&subCategory, id).Error
	if err != nil {
		return nil, err
	}
	return &subCategory, nil
}

func (r *subCategoryRepository) GetAll() ([]*SubCategory, error) {
	var subCategories []*SubCategory
	err := r.db.Where("estado = ?", true).Find(&subCategories).Error
	if err != nil {
		return nil, err
	}
	return subCategories, nil
}

func (r *subCategoryRepository) Update(subCategory *SubCategory) (string, error) {
	result := r.db.Model(&SubCategory{}).
		Where("id_subcategoria = ? AND estado = ?", subCategory.IdSubCategoria, true).
		Updates(subCategory)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_UPDATED, nil
}

func (r *subCategoryRepository) Delete(id uint) (string, error) {
	result := r.db.Model(&SubCategory{}).Where("id_subcategoria = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}
