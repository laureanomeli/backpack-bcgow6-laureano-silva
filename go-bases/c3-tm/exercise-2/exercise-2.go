package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, error := os.ReadFile("../exercise-1/products")
	if error != nil {
		panic(error)
	}
	fileAsString := fmt.Sprintf("%s", file)
	fmt.Print(fileAsString)
	reader := strings.NewReader(fileAsString)
	values, error := io.ReadAll(reader)
	if error != nil {
		panic(error)
	}
	for _, value := range values {
		fmt.Println(value)
	}
}
