package main

import "fmt"
import "errors"

func main() {
	const (
		perro = "perro"
		gato = "gato"
	)
	var cantidad float64
	animalPerro, err := Animal("perro")
	if(err != nil) {
		fmt.Println("Ocurrio un error")
	 } else{
		cantidad += animalPerro(5.0)
	 }
	animalGato, err := Animal("gato")
	if(err != nil) {
		fmt.Println("Ocurrio un error")
	 } else{
		cantidad += animalGato(8.0)
	}
	fmt.Println("La cantidad total a comprar es: " ,cantidad)

	
	
}

func animalPerro(cant float64) float64{
	return 10*cant
}
func animalGato(cant float64) float64{
	return 5*cant
}
func animalHamster(cant float64) float64{
	return 0.250*cant
}
func animalTarantula(cant float64) float64{
	return 0.150*cant
}


func  Animal(animal string) (func(cant float64) float64, error){
	switch animal {
	case "perro": 
		return animalPerro, nil
	case "gato": 
		return animalGato, nil
	case "hamster": 
		return animalHamster, nil
	case "tarantula":
		return animalTarantula, nil
	default:
		return nil, errors.New("No se paso un animal cargado")
 }
}