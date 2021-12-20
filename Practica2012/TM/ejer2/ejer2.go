package main

import	"fmt"
import "os"
import "encoding/csv"
import "strconv"

func main() {
	filename := "/Users/davividal/go/src/github.com/DavidVidalML/Clase1/Practica2012/TM/products.csv"
	fs1, _ := os.Open(filename)
	r1 := csv.NewReader(fs1)
	total := 0.0
	content, err := r1.ReadAll()
	if err != nil {
		fmt.Println("No se pudo leer el archivo")
	}
	for _, row := range content {
		fmt.Println("\t",row)
	}
	for i := 1; i < len(content); i++ {
			precio, _ := strconv.ParseFloat(content[i][1],64)	
			//fmt.Println("Precio es ", precio)		
			total += precio	
	}
	fmt.Println("El total es: ",total)
}