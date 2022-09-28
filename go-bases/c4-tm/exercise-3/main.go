package main

import (
	"fmt"
)

func main() {
	salary := 149000

	if salary < 150000 {
		error := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
		fmt.Println(error)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
