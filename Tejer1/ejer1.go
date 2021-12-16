package main

import "fmt"

func main() {
	var palabra string = "Jamaica"
	//var deletreada  =  strings.Split(palabra, "")
	//fmt.Println( "Cantidad de letras: ", len(palabra))
	//fmt.Printf("%q", deletreada)

	for _, element := range palabra {
		fmt.Println(string(element))
	}
}