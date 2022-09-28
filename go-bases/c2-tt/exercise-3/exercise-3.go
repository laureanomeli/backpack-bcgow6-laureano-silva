package main

type tienda struct {
	productos []Producto
}

type producto struct {
	tipo   string
	nombre string
	precio float64
}

type Producto interface {
	CalcularCosto()
}

func calcularcosto()

type Ecommerce interface {
	Total()
	Agregar()
}
