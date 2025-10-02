package bankaccount

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type BankAccountRepository interface {
	Create(bankAccount *BankAccount) error
	GetByID(id uint) (*BankAccount, error)
	GetAll() ([]*BankAccount, error)
	Update(bankAccount *BankAccount) error
	Delete(id uint) (string, error)
}

type bankAccountRepository struct {
	db *gorm.DB
}

func NewBankAccountRepository(db *gorm.DB) BankAccountRepository {
	return &bankAccountRepository{db}
}

func (r *bankAccountRepository) Create(bankAccount *BankAccount) error {
	return r.db.Create(bankAccount).Error
}

func (r *bankAccountRepository) GetByID(id uint) (*BankAccount, error) {
	var bankAccount BankAccount
	err := r.db.Where("estado = ?", true).First(&bankAccount, id).Error
	if err != nil {
		return nil, err
	}
	return &bankAccount, nil
}

func (r *bankAccountRepository) GetAll() ([]*BankAccount, error) {
	var bankAccounts []*BankAccount
	err := r.db.Where("estado = ?", true).Find(&bankAccounts).Error
	if err != nil {
		return nil, err
	}
	return bankAccounts, nil
}

func (r *bankAccountRepository) Update(bankAccount *BankAccount) error {
	return r.db.Model(&BankAccount{}).Where("id_cuenta_bancaria = ?", bankAccount.IdCuentaBancaria).Updates(bankAccount).Error
}

func (r *bankAccountRepository) Delete(id uint) (string, error) {
	result := r.db.Model(&BankAccount{}).Where("id_cuenta_bancaria = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}
