package main

import "fmt"


func main() {
	var mes int = 9
	switch mes {
	case 1:
		fmt.Println(mes,",Enero")
	case 2:
		fmt.Println(mes,",Febrero")
	case 3:
		fmt.Println(mes,",Marzo")
	case 4:
		fmt.Println(mes,",Abril")
	case 5:
		fmt.Println(mes,",Mayo")
	case 6:
		fmt.Println(mes,",Junio")
	case 7:
		fmt.Println(mes,",Julio")
	case 8:
		fmt.Println(mes,",Agosto")
	case 9:
		fmt.Println(mes,",Setiembre")
	case 10:
		fmt.Println(mes,",Octubre")
	case 11:
		fmt.Println(mes,",Noviembre")
	case 12:
		fmt.Println(mes,",Diciembre")
	}
}