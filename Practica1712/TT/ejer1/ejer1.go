package main

import "fmt"

func main() {
	e1:= Estudiante{
		Nombre:    "Mariano",
		Apellido:  "Bustamante",
		DNI: 39001395,
		FechaIng: Fecha{
			Dia:  1,
			Mes:  2,
			Anio: 1990,
		},
	}
	e1.details()
}
type Fecha struct {
	Dia, Mes, Anio int
}
type Estudiante struct {
	Nombre string
	Apellido string
	DNI int
	FechaIng Fecha
}

func (e *Estudiante)details() {
	fmt.Printf("\n\nE1: %+v, ", e)
}

