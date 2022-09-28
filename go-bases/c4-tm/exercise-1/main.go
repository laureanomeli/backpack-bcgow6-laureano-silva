package main

import "fmt"

type myError struct {
	message string
}

func (e *myError) Error() string {
	return fmt.Sprint("error: el salario ingresado no alcanza el mínimo imponible")
}

func main() {
	salary := 159000

	if salary < 150000 {
		var error *myError
		fmt.Printf("%v \n", error)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
