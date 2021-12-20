package main

import	"fmt"
import "os"
import "encoding/csv"
/*
Una empresa que se encarga de vender productos de limpieza necesita: 
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.


**/
func main() {


 data := [][] string{
	 {"id","precio","cantidad"},
	 {"1","300","10"},
	 {"2","450","5"},
	 {"10","25","200"},
 }
 //Crear CSV
 csvFile, err := os.Create("products.csv")
 	if err != nil {
		fmt.Println("Error")
	}

csvwriter := csv.NewWriter(csvFile)
 
for _, prodRow := range data {
	_ = csvwriter.Write(prodRow)
}

csvwriter.Flush()
csvFile.Close()
}