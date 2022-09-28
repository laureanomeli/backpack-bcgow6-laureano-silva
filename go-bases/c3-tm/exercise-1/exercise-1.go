package main

import "fmt"

type product struct {
	id       int
	price    float64
	quantity int
}

func main() {
	products := []product{[1, 100, 50, 2, 200, 150, 3, 400, 510}
	//productsToBytes := byte(products)
	fmt.Print(products)
	/* err := os.WriteFile("./products", products, 0644)
	if err != nil {
		print(err)
	} */
}
