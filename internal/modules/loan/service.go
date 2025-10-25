package loan

import (
	"errors"
	"pistolistoBE/internal/modules/avaluo"
	"pistolistoBE/internal/modules/contrato"
)

type LoanService interface {
	CreateLoanFromContrato(idContrato uint) (*Loan, error)
	GetByID(id uint) (*Loan, error)
	GetAll() ([]*Loan, error)
	Update(id uint, loan *Loan) error
	Delete(id uint) error
}

type loanService struct {
	repo        LoanRepository
	contratoSrv contrato.ContratoService
	avaluoSrv   avaluo.AvaluoService
}

func NewLoanService(repo LoanRepository, contratoSrv contrato.ContratoService, avaluoSrv avaluo.AvaluoService) LoanService {
	return &loanService{repo, contratoSrv, avaluoSrv}
}

func (s *loanService) CreateLoanFromContrato(idContrato uint) (*Loan, error) {
	// 1️⃣ Obtener contrato
	contrato, err := s.contratoSrv.GetByID(idContrato)
	if err != nil {
		return nil, errors.New("no se encontró el contrato asociado")
	}

	// 2️⃣ Obtener avalúo desde contrato
	avaluo, err := s.avaluoSrv.GetByID(uint(contrato.IdAvaluo))
	if err != nil {
		return nil, errors.New("no se encontró el avalúo asociado al contrato")
	}

	// 3️⃣ Crear préstamo usando el monto avaluado
	loan := &Loan{
		IdContrato:        int(idContrato),
		MontoPrestado:     avaluo.MontoAvaluado,
		CobroDespuesXDias: new(int),
	}
	*loan.CobroDespuesXDias = 30 // mock: cobro cada 30 días

	err = s.repo.Create(loan)
	if err != nil {
		return nil, err
	}

	return loan, nil
}

func (s *loanService) GetByID(id uint) (*Loan, error) {
	return s.repo.GetByID(id)
}

func (s *loanService) GetAll() ([]*Loan, error) {
	return s.repo.GetAll()
}

func (s *loanService) Update(id uint, loan *Loan) error {
	return s.repo.Update(id, loan)
}

func (s *loanService) Delete(id uint) error {
	return s.repo.Delete(id)
}
