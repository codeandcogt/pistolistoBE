package loan

import "gorm.io/gorm"

type LoanRepository interface {
	Create(loan *Loan) error
	GetByID(id uint) (*Loan, error)
	GetAll() ([]*Loan, error)
	Update(id uint, loan *Loan) error
	Delete(id uint) error
}

type loanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanRepository{db}
}

func (r *loanRepository) Create(loan *Loan) error {
	return r.db.Create(loan).Error
}

func (r *loanRepository) GetByID(id uint) (*Loan, error) {
	var loan Loan
	err := r.db.Where("id_prestamo = ? AND estado = ?", id, true).First(&loan).Error
	if err != nil {
		return nil, err
	}
	return &loan, nil
}

func (r *loanRepository) GetAll() ([]*Loan, error) {
	var loans []*Loan
	err := r.db.Where("estado = ?", true).Order("fecha_creacion DESC").Find(&loans).Error
	if err != nil {
		return nil, err
	}
	return loans, nil
}

func (r *loanRepository) Update(id uint, loan *Loan) error {
	return r.db.Model(&Loan{}).Where("id_prestamo = ? AND estado = ?", id, true).Updates(loan).Error
}

func (r *loanRepository) Delete(id uint) error {
	return r.db.Model(&Loan{}).Where("id_prestamo = ? AND estado = ?", id, true).Update("estado", false).Error
}
