package calculadora

import (
	"fmt"
	"sort"
)

func Restar(num1, num2 int) int {
	return num1 - num2
}

func Ordenar(numbers []int) []int {
	sort.Slice(numbers, func(i, q int) bool {
		return numbers[i] < numbers[q]
	})
	return numbers
}

func Dividir(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("no se puede dividir por 0")
	}
	return num1 / num2, nil
}
