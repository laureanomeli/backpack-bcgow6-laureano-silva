package main

import "fmt"

func calculateSalary(minutes int, category string) float64 {
	salaryA := 3000.0 / 60.0
	salaryB := 1500.0 / 60.0
	salaryC := 1000.0 / 60.0
	salary := .0
	switch category {
	case "c":
		salary = salaryC * float64(minutes)
	case "b":
		salary = salaryB * float64(minutes) * 1.2
	case "a":
		salary = salaryA * float64(minutes) * 1.5
	}
	return salary
}

func main() {
	fmt.Println(calculateSalary(60, "a"))
	fmt.Println(calculateSalary(65, "a"))
	fmt.Println(calculateSalary(60, "b"))
	fmt.Println(calculateSalary(65, "b"))
	fmt.Println(calculateSalary(60, "c"))
	fmt.Println(calculateSalary(65, "c"))
}
