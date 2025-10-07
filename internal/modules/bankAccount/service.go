package bankAccount

type BankAccountService interface {
	Create(bankAccount *BankAccount) error
	GetByID(id uint) (*BankAccount, error)
	GetAll() ([]*BankAccount, error)
	Update(bankAccount *BankAccount) (string, error)
	Delete(id uint) (string, error)
}

type bankAccountService struct {
	repo BankAccountRepository
}

func NewBankAccountService(repo BankAccountRepository) BankAccountService {
	return &bankAccountService{repo}
}

func (s *bankAccountService) Create(bankAccount *BankAccount) error {
	return s.repo.Create(bankAccount)
}

func (s *bankAccountService) GetByID(id uint) (*BankAccount, error) {
	return s.repo.GetByID(id)
}

func (s *bankAccountService) GetAll() ([]*BankAccount, error) {
	return s.repo.GetAll()
}

func (s *bankAccountService) Update(bankAccount *BankAccount) (string, error) {
	return s.repo.Update(bankAccount)
}

func (s *bankAccountService) Delete(id uint) (string, error) {
	return s.repo.Delete(id)
}
