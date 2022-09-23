package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("Benjamin age is: ", employees["Benjamin"])

	over21 := 0
	for _, employee := range employees {
		if employee > 21 {
			over21++
		}
	}

	fmt.Println("employes over 21 years old: ", over21)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
}
