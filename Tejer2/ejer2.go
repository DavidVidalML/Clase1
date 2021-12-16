package main

import "fmt"


func main() {
	var precio float32 = 200
	var porc float32 = 30
	var desc float32 = porc/100 
	fmt.Println("El precio del producto con el descuento obtenido es ", precio -(precio*desc))

}