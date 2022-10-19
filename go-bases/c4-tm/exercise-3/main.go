package main

import (
	"fmt"
	"os"
)

func main() {
	salary := 149000

	if salary < 150000 {
		err := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
