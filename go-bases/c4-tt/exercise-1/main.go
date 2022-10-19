package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "./main.go"

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover de este panic: ", err)
		}
		fmt.Println("Ejecucion finalizada")
	}()

	file, err := os.Open(filePath)
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}

	fmt.Println(file)
}
