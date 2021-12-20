package main

import "fmt"


func main() {
	sueldo := calcularSal(4000,"A")
	fmt.Println("El sueldo del empleado es: ", sueldo )
}

func  calcularSal(minutos float64, cat string) float64 {
	sueldo := 0.0
	horas := minutos/60.0
	switch cat {
	case "C": 
		sueldo = horas *1000
	case "B":
		sueldo = horas * 1500 
		sueldo += sueldo*0.2
	case "A":
		sueldo = horas * 3000 
		sueldo += sueldo*0.5	
	}
	return sueldo
}