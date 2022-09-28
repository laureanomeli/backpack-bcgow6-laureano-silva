package main

import (
	"errors"
	"fmt"
)

func getMedia(values ...float64) (float64, error) {
	sum, count := .0, .0
	for _, value := range values {
		if value < 0 {
			return -1, errors.New("can`t be negative values")
		} else {
			sum += value
			count++
		}
	}
	return sum / count, nil
}

func main() {
	fmt.Println(getMedia(4, 5, 6, 7, 3, 2, 3))
	fmt.Println(getMedia(-1, 2, 3))
}
