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
	GetByCodigo(codigo string) (*Banco, error)
	GetByTipo(tipo string) ([]*Banco, error)
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
	err := r.db.First(&banco, id).Error
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
	return r.db.Model(&Banco{}).Where("id_banco = ?", id).Updates(banco).Error
}

func (r *bancoRepository) Delete(id uint) error {
	return r.db.Model(&Banco{}).Where("id_banco = ?", id).Update("estado", false).Error
}

func (r *bancoRepository) GetByCodigo(codigo string) (*Banco, error) {
	var banco Banco
	err := r.db.Where("codigo_banco = ? AND estado = ?", codigo, true).First(&banco).Error
	if err != nil {
		return nil, err
	}
	return &banco, nil
}

func (r *bancoRepository) GetByTipo(tipo string) ([]*Banco, error) {
	var bancos []*Banco
	err := r.db.Where("tipo_banco = ? AND estado = ?", tipo, true).Find(&bancos).Error
	if err != nil {
		return nil, err
	}
	return bancos, nil
}
