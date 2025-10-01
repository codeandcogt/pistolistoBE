package subsidiary

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type SubsidiaryRepository interface {
	Create(subsidiary *Subsidiary) error
	GetByID(id uint) (*Subsidiary, error)
	GetAll() ([]*Subsidiary, error)
	Update(subsidiary *Subsidiary) error
	Delete(id uint) (string, error)
}

type subsidiaryRepository struct {
	db *gorm.DB
}

func NewSubsidiaryRepository(db *gorm.DB) SubsidiaryRepository {
	return &subsidiaryRepository{db}
}

func (r *subsidiaryRepository) Create(subsidiary *Subsidiary) error {
	return r.db.Create(subsidiary).Error
}

func (r *subsidiaryRepository) GetByID(id uint) (*Subsidiary, error) {
	var subsidiary Subsidiary
	err := r.db.Where("estado = ?", true).First(&subsidiary, id).Error
	if err != nil {
		return nil, err
	}
	return &subsidiary, nil
}

func (r *subsidiaryRepository) GetAll() ([]*Subsidiary, error) {
	var subsidiaries []*Subsidiary
	err := r.db.Where("estado = ?", true).Find(&subsidiaries).Error
	if err != nil {
		return nil, err
	}
	return subsidiaries, nil
}

func (r *subsidiaryRepository) Update(subsidiary *Subsidiary) error {
	return r.db.Model(&Subsidiary{}).Where("id_sucursal = ?", subsidiary.IdSucursal).Updates(subsidiary).Error
}

func (r *subsidiaryRepository) Delete(id uint) (string, error) {
	result := r.db.Model(&Subsidiary{}).Where("id_sucursal = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}
