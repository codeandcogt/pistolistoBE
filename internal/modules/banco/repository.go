package banco

import (
	"gorm.io/gorm"
)

type BancoRepository interface {
	Create(banco *Banco) error
	GetByID(id uint) (*Banco, error)
	GetAll() ([]*Banco, error)
	Update(id uint, banco *Banco) error
	Delete(id uint) error
}

type bancoRepository struct {
	db *gorm.DB
}

func NewBancoRepository(db *gorm.DB) BancoRepository {
	return &bancoRepository{db}
}

func (r *bancoRepository) Create(banco *Banco) error {
	return r.db.Create(banco).Error
}

func (r *bancoRepository) GetByID(id uint) (*Banco, error) {
	var banco Banco
	err := r.db.Where("id_banco = ? AND estado = ?", id, true).First(&banco).Error
	if err != nil {
		return nil, err
	}
	return &banco, nil
}

func (r *bancoRepository) GetAll() ([]*Banco, error) {
	var bancos []*Banco
	err := r.db.Where("estado = ?", true).Find(&bancos).Error
	if err != nil {
		return nil, err
	}
	return bancos, nil
}

func (r *bancoRepository) Update(id uint, banco *Banco) error {
	return r.db.Model(&Banco{}).Where("id_banco = ? AND estado = ?", id, true).Updates(banco).Error
}

func (r *bancoRepository) Delete(id uint) error {
	return r.db.Model(&Banco{}).Where("id_banco = ? AND estado = ?", id, true).Update("estado", false).Error
}
