package main

import "fmt"

type Alumno struct {
	nombre   string
	apellido string
	dni      int
	fecha    string
}

func (a Alumno) detalle() {
	fmt.Printf("Nombre: %s \n", a.nombre)
	fmt.Printf("Apellido: %s \n", a.apellido)
	fmt.Printf("DNI: %d \n", a.dni)
	fmt.Printf("Fecha: %s \n", a.fecha)
}

func main() {
	alumno1 := Alumno{"Laureano", "Silva", 1111111, "1/1/2022"}

	alumno1.detalle()
}
