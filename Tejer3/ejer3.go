package main

import "fmt"


func main() {
	var edad int = 23
	var empleado bool = true
	var antig int = 1
	var sueldo float64 = 60000
	if(edad>22 && empleado == true && antig>1){
		if(sueldo>100000){
			fmt.Println("No le cobrara interes")
		}else{
			fmt.Println("Le cobrara interes")
		}
	}else{
		fmt.Println("No le otorgara prestamo")
	}
}