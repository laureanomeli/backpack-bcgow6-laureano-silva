package main

import (
	"errors"
	"fmt"
)

func calculateSalary(hours, rateByHour float64) (salary float64, err error) {
	if hours < 80 {
		err = errors.New("error: el trabajoador no puede haber trabjado menos de 80 hs mensuales")
		return
	}
	salary = hours * rateByHour
	if salary > 150000 {
		salary = salary - salary*.1
	}
	return
}

func calculateBonus(last12Salaries ...float64) (bonus float64, err error) {
	maxSalary := 0.
	workedMonths := 0
	for _, salary := range last12Salaries {
		if salary < 0 {
			err = errors.New("salary cannot be negative")
		}
		if salary > maxSalary {
			maxSalary = salary
		}
		workedMonths++
	}
	bonus = maxSalary / 12 * float64(workedMonths)
	return
}

func main() {
	hoursByMonth := []float64{75, 90, 120}
	rate := 3000.
	salaries := []float64{}

	for _, hours := range hoursByMonth {
		salary, err := calculateSalary(hours, rate)
		if err != nil {
			fmt.Print(err.Error())
		}
		salaries = append(salaries, salary)
	}
	fmt.Println(salaries)
	bonus, err := calculateBonus(salaries...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bonus)
}
