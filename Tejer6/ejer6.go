package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("Edad de Benjamin: ",employees["Benjamin"])

	//Contar mayores de 21 años
	var mayor int 
	for  _,element := range employees {
		if(element>21){
			mayor++
		}
	}
	fmt.Println("Mayores de 21 : ", mayor)

	//Agregar Federico 25 años
	employees["Federico"] = 25
	fmt.Println(employees)
	
	//Eliminar Pedro
	delete(employees, "Pedro")
	fmt.Println(employees)
}