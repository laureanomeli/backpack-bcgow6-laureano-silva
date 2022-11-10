package transactions

import (
	"fmt"

	"github.com/laureanomeli/backpack-bcgow6-laureano-silva/go-web/c2-tt/pkg/store"
)

type Transaction struct {
	Id       int     `json:"id"`
	Code     string  `json:"code"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Emitter  string  `json:"emitter"`
	Receiver string  `json:"receiver"`
	Date     string  `json:"date"`
}

// ***Importado**//
type Repository interface {
	GetAll() ([]Transaction, error)
	Store(id int, code, currency string, amount float64, emitter, receiver, date string) (Transaction, error)
	LastID() (int, error)
	Update(id int, code, currency string, amount float64, emitter, receiver, date string) (Transaction, error)
	Delete(id int) error
	UpdateCodeAndAmount(id int, code string, amount float64) (Transaction, error)
}

type repository struct {
	db store.Store
} //struct implementa los metodos de la interfaz

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(id int, code, currency string, amount float64, emitter, receiver, date string) (Transaction, error) {
	newTransaction := Transaction{id, code, currency, amount, emitter, receiver, date}
	var transactions []Transaction
	if err := r.db.Read(&transactions); err != nil {
		return Transaction{}, err
	}
	transactions = append(transactions, newTransaction)
	if err := r.db.Write(transactions); err != nil {
		return Transaction{}, err
	}
	return newTransaction, nil
}

func (r *repository) GetAll() ([]Transaction, error) {
	var transactions []Transaction
	if err := r.db.Read(&transactions); err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) LastID() (int, error) {
	var transactions []Transaction
	if err := r.db.Read(&transactions); err != nil {
		return 0, err
	}
	if len(transactions) == 0 {
		return 0, nil
	}

	return transactions[len(transactions)-1].Id, nil
}

func (r *repository) Update(id int, code, currency string, amount float64, emitter, receiver, date string) (Transaction, error) {
	transaction := Transaction{Code: code, Currency: currency, Amount: amount, Emitter: emitter, Receiver: receiver, Date: date}
	updated := false
	var transactions []Transaction
	if err := r.db.Read(&transactions); err != nil {
		return Transaction{}, err
	}
	for i := range transactions {
		if transactions[i].Id == id {
			transaction.Id = id
			transactions[i] = transaction
			updated = true
		}
	}
	if !updated {
		return Transaction{}, fmt.Errorf("Transaction %d not found", id)
	}
	return transaction, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	var transactions []Transaction
	if err := r.db.Read(&transactions); err != nil {
		return err
	}
	for i := range transactions {
		if transactions[i].Id == id {
			index = id
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Transaction %d not found", id)
	}
	transactions = append(transactions[:index], transactions[index+1:]...)
	return nil
}

func (r *repository) UpdateCodeAndAmount(id int, code string, amount float64) (Transaction, error) {
	var transaction Transaction
	updated := false
	var transactions []Transaction
	if err := r.db.Read(&transactions); err != nil {
		return transaction, err
	}
	for index := range transactions {
		if transactions[index].Id == id {
			transactions[index].Code = code
			transactions[index].Amount = amount
			transaction = transactions[index]
			updated = true
		}
	}
	if !updated {
		return Transaction{}, fmt.Errorf("Transaction %d not found", id)
	}
	return transaction, nil
}
