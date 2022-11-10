package transactions

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/laureanomeli/backpack-bcgow6-laureano-silva/testing/c2-tt/pkg/store"
	"github.com/stretchr/testify/assert"
)

/*
Ejercicio 2 - Test Unitario UpdateName()
Diseñar Test de UpdateName, donde se valide que la respuesta retornada sea correcta para la actualización del nombre de un producto específico. Y además se compruebe que efectivamente se usa el método “Read” del Storage para buscar el producto. Para esto:
    1. Crear un mock de Storage, dicho mock debe contener en su data un producto específico cuyo nombre puede ser “Before Update”.
    2. El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado. Puede ser a través de un boolean como se observó en la clase.
    3. Para dar el test como OK debe validarse que al invocar el método del Repository UpdateName, con el id del producto mockeado y con el nuevo nombre “After Update”, efectivamente haga la actualización. También debe validarse que el método Read haya sido ejecutado durante el test.
*/

/* Test with Mock */
func TestUpdate1(t *testing.T) {
	id, code, currency, amount, emitter, receiver, date := 1, "Update After", "usd", 100., "lala", "lele", "hoy"
	transactions := []Transaction{{Id: 1, Code: "Update Before", Currency: "ars", Amount: 1., Emitter: "tata", Receiver: "tete", Date: "mañana"}}

	data, _ := json.Marshal(transactions)
	mock := store.Mock{Data: data}

	stubStore := store.FileStore{
		FilePath: "",
		Mock:     &mock,
	}

	r := NewRepository(&stubStore)
	transactionUpdated, err := r.Update(id, code, currency, amount, emitter, receiver, date)
	assert.Nil(t, err)

	assert.True(t, mock.ReadInvoked)
	assert.Equal(t, id, transactionUpdated.Id)
	assert.Equal(t, code, transactionUpdated.Code)
}

func TestUpdateTransactionNotExists(t *testing.T) {
	id, code, currency, amount, emitter, receiver, date := 100, "Update After", "usd", 100., "lala", "lele", "hoy"
	transactions := []Transaction{{Id: 1, Code: "Update Before", Currency: "ars", Amount: 1., Emitter: "tata", Receiver: "tete", Date: "mañana"}}

	data, _ := json.Marshal(transactions)
	mock := store.Mock{Data: data}

	stubStore := store.FileStore{
		FilePath: "",
		Mock:     &mock,
	}

	r := NewRepository(&stubStore)
	transactionUpdated, err := r.Update(id, code, currency, amount, emitter, receiver, date)
	assert.NotNil(t, err)
	assert.Equal(t, transactionUpdated, Transaction{})
}

func TestUpdateTransactionFileUnavailable(t *testing.T) {
	id, code, currency, amount, emitter, receiver, date := 100, "Update After", "usd", 100., "lala", "lele", "hoy"
	transactions := fmt.Errorf("")

	data, _ := json.Marshal(transactions)
	mock := store.Mock{Data: data}

	stubStore := store.FileStore{
		FilePath: "",
		Mock:     &mock,
	}

	r := NewRepository(&stubStore)
	transaction, err := r.Update(id, code, currency, amount, emitter, receiver, date)
	assert.NotNil(t, err)
	assert.Equal(t, transaction, Transaction{})
}

func TestUpdateCodeAndAmount(t *testing.T) {
	id, code, amount := 1, "Update After", 100.
	transactions := []Transaction{{Id: 1, Code: "Update Before", Amount: 1.}}

	data, _ := json.Marshal(transactions)
	mock := store.Mock{Data: data}

	stubStore := store.FileStore{
		FilePath: "",
		Mock:     &mock,
	}

	r := NewRepository(&stubStore)
	transactionUpdated, err := r.UpdateCodeAndAmount(id, code, amount)
	assert.Nil(t, err)

	assert.True(t, mock.ReadInvoked)
	assert.Equal(t, id, transactionUpdated.Id)
	assert.Equal(t, code, transactionUpdated.Code)

	emptyTransaction, err1 := r.UpdateCodeAndAmount(100, "cualquiera", 100.)
	assert.NotNil(t, err1)
	assert.Equal(t, emptyTransaction, Transaction{})
}

func TestUpdateCodeAndAmountTransactionNotExists(t *testing.T) {
	id, code, amount := 100, "Update After", 100.
	transactions := []Transaction{{Id: 1, Code: "Update Before", Amount: 1.}}

	data, _ := json.Marshal(transactions)
	mock := store.Mock{Data: data}

	stubStore := store.FileStore{
		FilePath: "",
		Mock:     &mock,
	}

	r := NewRepository(&stubStore)
	transactionUpdated, err := r.UpdateCodeAndAmount(id, code, amount)
	assert.NotNil(t, err)
	assert.Equal(t, transactionUpdated, Transaction{})
}

func TestUpdateCodeAndAmountTransactionFileUnavailable(t *testing.T) {
	transactions := fmt.Errorf("")

	data, _ := json.Marshal(transactions)
	mock := store.Mock{Data: data}

	stubStore := store.FileStore{
		FilePath: "",
		Mock:     &mock,
	}

	r := NewRepository(&stubStore)
	transaction, err := r.UpdateCodeAndAmount(100, "algo", 100.)
	assert.NotNil(t, err)
	assert.Equal(t, transaction, Transaction{})
}

func TestDeleteTransactionNotExists(t *testing.T) {
	transactions := []Transaction{{Id: 1, Code: "cfvxd", Amount: 1.}}

	data, _ := json.Marshal(transactions)
	mock := store.Mock{Data: data}

	stubStore := store.FileStore{
		FilePath: "",
		Mock:     &mock,
	}

	r := NewRepository(&stubStore)
	err := r.Delete(100)
	assert.NotNil(t, err)
}

func TestDeleteTransactionFileUnavailable(t *testing.T) {
	transactions := fmt.Errorf("")

	data, _ := json.Marshal(transactions)
	mock := store.Mock{Data: data}

	stubStore := store.FileStore{
		FilePath: "",
		Mock:     &mock,
	}

	r := NewRepository(&stubStore)
	err := r.Delete(100)
	assert.NotNil(t, err)
}

/*
Ejercicio 1 - Test Unitario GetAll()
Generar un Stub del Store cuya función “Read” retorne dos productos con las especificaciones que deseen. Comprobar que GetAll() retorne la información exactamente igual a la esperada. Para esto:
    1. Dentro de la carpeta /internal/products, crear un archivo repository_test.go con el test diseñado.
*/

/* Test with Stub */
func TestGetAll(t *testing.T) {
	transactions := []Transaction{
		{
			Id:     1,
			Code:   "un codigo",
			Amount: 12,
		},
		{
			Id:     2,
			Code:   "otro codigo",
			Amount: 1,
		},
	}

	data, _ := json.Marshal(transactions)
	dbMock := store.Mock{
		Data:  data,
		Error: nil,
	}

	stubStore := store.FileStore{
		FilePath: "",
		Mock:     &dbMock,
	}

	r := NewRepository(&stubStore)
	transactionsGet, err := r.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, transactionsGet, transactions)
}

func TestGetAllError(t *testing.T) {
	errorExpected := errors.New("error for GetAll")
	dbMock := store.Mock{ // No le pasamos información
		Error: errorExpected, // Deberia fallar
	}

	stubStore := store.FileStore{
		FilePath: "",
		Mock:     &dbMock,
	}

	r := NewRepository(&stubStore)
	transactions, err := r.GetAll()

	assert.Equal(t, errorExpected, err)
	assert.Nil(t, transactions)
}
