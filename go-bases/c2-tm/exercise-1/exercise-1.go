package main

import "fmt"

func getTaxes(wage float64) float64 {
	tax := .0
	if wage > 50000 {
		tax = wage * .17
	}
	if wage > 150000 {
		tax = tax + wage*.1
	}
	return tax
}

func main() {
	fmt.Println(getTaxes(50000))
	fmt.Println(getTaxes(51000))
	fmt.Println(getTaxes(150000))
	fmt.Println(getTaxes(151000))
}
