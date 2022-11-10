package transactions

import (
	"encoding/json"
	"testing"

	"github.com/laureanomeli/backpack-bcgow6-laureano-silva/testing/c2-tt/pkg/store"
	"github.com/stretchr/testify/assert"
)

/*Ejercicio 1 - Service/Repo/Db Update()

Diseñar un test que pruebe en la capa service, el método o función Update(). Para lograrlo se deberá:
	1. Crear un mock de Storage, dicho mock debe contener en su data un producto con las especificaciones que desee.
	2. El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado.
	3. Para dar el test como OK debe validarse que al invocar el método del Service Update(),  retorne el producto con
	mismo Id y los datos actualizados. Validar también que  Read() del Store haya sido ejecutado durante el test.
*/

func TestUpdate(t *testing.T) {
	transactions := []Transaction{
		{
			Id:       1,
			Code:     "lala",
			Currency: "usd",
			Amount:   100,
			Emitter:  "lnjlk",
			Receiver: "ablib",
			Date:     "ohm",
		},
		{
			Id:       2,
			Code:     "lele",
			Currency: "ars",
			Amount:   1000,
			Emitter:  "lnjlk",
			Receiver: "ablib",
			Date:     "oh",
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

	expTransaction := Transaction{
		Id:       1,
		Code:     "2342",
		Currency: "lolo",
		Amount:   100,
		Emitter:  "lala",
		Receiver: "lele",
		Date:     "123",
	}

	repository := NewRepository(&stubStore)
	service := NewService(repository)
	transaction, err := service.Update(1, "2342", "lolo", 100, "lala", "lele", "123")
	assert.Nil(t, err)
	assert.Equal(t, expTransaction, transaction)
}

/*Ejercicio 2 - Service/Repo/Db Delete()
Diseñar un test que pruebe en la capa service, el método o función Delete(). Se debe probar la
correcta eliminación de un producto, y el error cuando el producto no existe. Para lograrlo puede:
	1. Crear un mock de Storage, dicho mock debe contener en su data un producto con las especificaciones que desee.
	2. Ejecutar el test con dos id’s de producto distintos, siendo uno de ellos un id inexistente en el Mock de Storage.
	3. Para dar el test como OK debe validarse que efectivamente el producto borrado ya no exista en Storage luego del Delete().
	También que cuando se intenta borrar un producto  inexistente, se debe obtener el error correspondiente.
*/

func TestDelete(t *testing.T) {
	transactions := []*Transaction{
		{
			Id:       1,
			Code:     "lala",
			Currency: "usd",
			Amount:   100,
			Emitter:  "lnjlk",
			Receiver: "ablib",
			Date:     "oh",
		},
		{
			Id:       2,
			Code:     "lele",
			Currency: "ars",
			Amount:   1000,
			Emitter:  "lnjlk",
			Receiver: "ablib",
			Date:     "oh",
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

	repository := NewRepository(&stubStore)
	service := NewService(repository)
	err := service.Delete(1)

	assert.Nil(t, err)
}
