package main

import (
	"errors"
	"fmt"
)

func main() {
	salary := 149000

	if salary < 150000 {
		fmt.Print(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
