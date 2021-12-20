package main

import "fmt"


func main() {
	// estudiantes := make([]string, 13) append lo que hace es generar un nuevo array £
	estudiantes := []string { "Benjamin","Nahuel", "Brenda", "Marcos","Pedro", "Axel", "Alez","Dolores", "Federico", "Hernán","Leandro", "Eduardo", "Duvraschka"}
	fmt.Println(estudiantes[:])
	fmt.Printf("%d\n", cap(estudiantes))
	estudiantes=append(estudiantes,"Gabriela")
	fmt.Println(estudiantes[:])
	fmt.Printf("%d\n", cap(estudiantes))
}