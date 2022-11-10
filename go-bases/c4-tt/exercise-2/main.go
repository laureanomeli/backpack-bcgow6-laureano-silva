package main

import (
	"fmt"
	"os"
)

type Client struct {
	legajo    int
	nombre    string
	apellido  string
	dni       int
	telefono  int
	domicilio string
}

var counter int = 0

func (client *Client) generateRecordNumber(counter int) (number int) {
	counter++
	if counter%2 != 0 {
		client.legajo = counter
		return counter
	} else {
		panic("number is odd")
	}
}

func readFile(filePath string) (fileContent []byte, err error) {
	fileContent, err = os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	cliente1 := Client{}
	cliente1.generateRecordNumber(counter)
	cliente1.nombre = "a"
	cliente1.apellido = "b"
	cliente1.dni = 22323
	cliente1.telefono = 78567
	cliente1.domicilio = "lala"

	fmt.Println(cliente1)
}
