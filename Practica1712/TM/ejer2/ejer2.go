package main

import "fmt"
import "errors"


func main() {
promedio,error:= calcularProm()
	if error != nil {
		fmt.Println("El alumno no tiene notas")
	}else{
		fmt.Println("El promedio de sus notas es: ", promedio)
	}
}

func  calcularProm(notas ...int) (int,error) {
	 total := 0
	if(len(notas) == 0){
		return 0,errors.New("El alumno no tiene notas")
	}else{
		for _, nota := range notas {
			total += nota
		}
		prom := total/ len(notas)
		return prom,nil
	}
}