package main

import "fmt"
import "errors"

func main() {
	const (
		minimo = "minimo"
		promedio = "promedio"
		maximo = "maximo"
	 )
	 minFunc, err := operacion(minimo)
	 if(err != nil) {
		fmt.Println("Ocurrio un error")
	 } else{
		valorMinimo := minFunc(2,3,3,4,1,2,4,5)
		fmt.Println("El minimo es: ",valorMinimo)
	 }
	 promFunc, err := operacion(promedio)
	 if(err != nil) {
		fmt.Println("Ocurrio un error")
	 } else{
		valorPromedio := promFunc(2,3,3,4,1,2,4,5)
		fmt.Println("El promedio es: ",valorPromedio)
	 }
	 maxFunc, err := operacion(maximo)
	 if(err != nil) {
		fmt.Println("Ocurrio un error")
	 } else{
		valorMaximo := maxFunc(2,3,3,4,1,2,4,5)
		fmt.Println("El maximo es: ",valorMaximo)
	 }

}
func minFunc(values ...float64) (float64) {
	menor := values[0]
	for _, value := range values {
		if(value < menor) {
			menor= value
		}
	}
	return menor
}
func maxFunc(values ...float64) float64 {
	mayor := values[0]
	for _, value := range values {
		if(value > mayor) {
			mayor= value
		}
	}
	return mayor
}
func promFunc(values ...float64) float64 {
	total := 0.0
	prom := 0.0
	for _, value := range values {
		total += value
	}
	prom = total/float64(len(values))
	return prom
}


func  operacion(estadistica string) (func(values ...float64) float64, error){

	switch estadistica {
	case "minimo":
		return minFunc,nil
	case "maximo":
		return maxFunc,nil
	case "promedio":
		return promFunc,nil
	default:
		return nil, errors.New("No se paso una operacion correcta")

	}
}