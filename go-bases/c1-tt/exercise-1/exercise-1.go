package main

import "fmt"

func main() {
	word := "chuncaco"
	fmt.Println("number of letters: ", len(word))
	for _, character := range word {
		fmt.Println(string(character))
	}
}
