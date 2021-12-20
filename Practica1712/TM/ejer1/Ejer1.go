package main

import "fmt"


func main() {
	impu := calcularImp(60000)
	fmt.Println("Se le descontara: ", impu)
}

func calcularImp(sueldo float64) float64 {
	var impuesto float64 = 0.0
	if(sueldo>50000){
		if(sueldo>150000){
			impuesto = sueldo*0.10
		}
		impuesto += sueldo*0.17
	}
	return impuesto
}