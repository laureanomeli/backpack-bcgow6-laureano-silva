package transactions

type Service interface {
	GetAll() ([]Transaction, error)
	Store(code, currency string, amount float64, emitter, receiver, date string) (Transaction, error)
	Update(id int, code, currency string, amount float64, emitter, receiver, date string) (Transaction, error)
	Delete(id int) error
	UpdateCodeAndAmount(id int, code string, amount float64) (Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Transaction, error) {
	transaction, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *service) Store(code, currency string, amount float64, emitter, receiver, date string) (Transaction, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Transaction{}, err
	}

	lastID++

	transaction, err := s.repository.Store(lastID, code, currency, amount, emitter, receiver, date)
	if err != nil {
		return Transaction{}, err
	}

	return transaction, nil
}

func (s *service) Update(id int, code, currency string, amount float64, emitter, receiver, date string) (Transaction, error) {

	return s.repository.Update(id, code, currency, amount, emitter, receiver, date)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateCodeAndAmount(id int, code string, amount float64) (Transaction, error) {
	return s.repository.UpdateCodeAndAmount(id, code, amount)
}
