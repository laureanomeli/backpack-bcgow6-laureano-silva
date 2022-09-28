package main

import "fmt"

type Matrix struct {
	values    []float64
	heigth    int
	width     int
	cuadratic bool
	maxValue  float64
}

func (m *Matrix) Set(input ...float64) {
	m.values = input

	for _, value := range input {
		if value > m.maxValue {
			m.maxValue = value
		}
	}
}

func (m Matrix) Print() {
	for i := 0; i < len(m.values); i++ {
		if (i+1)%m.width == 0 {
			fmt.Println(m.values[i])
		} else {
			fmt.Print(m.values[i], " ")
		}
	}
}

func main() {
	myMatrix := Matrix{heigth: 3, width: 4}

	myMatrix.Set(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)

	myMatrix.Print()
}
