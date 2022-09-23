package main

import "fmt"

func main() {
	age := 22
	isEmployee := true
	laborOld := 5
	wage := 120000
	if age <= 22 || isEmployee == false || laborOld < 1 {
		fmt.Println("You can not acces this loan by now.")
	} else if wage < 100000 {
		fmt.Println("you must pay interest for this loan")
	} else {
		fmt.Println("You can request a loan with 0% interest!")
	}
}
