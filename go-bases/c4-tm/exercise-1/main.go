package main

import "fmt"

type myError struct {
}

func (e *myError) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

func main() {
	salary := 149000

	if salary < 150000 {
		var error *myError
		fmt.Printf("%v \n", error)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
