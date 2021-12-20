package main

import "fmt"

func main() {
	a := []float64{3, 4, 7}
    b := []float64{3, 4, 6}
	matrix := [][]float64{a, b}
	m := Matrix{}
	m.set(matrix)
	m.print()
}

type Matrix struct {
	DimAlto int
	DimAncho int
	Cuadratica bool
	MaxValue float64
	Values [][] float64
}

func (m *Matrix)set(matrix [][] float64) {
	m.Values = matrix
	m.DimAncho = len(matrix[0])
	m.DimAlto = len(matrix)
	m.Cuadratica = (m.DimAncho == m.DimAlto)
	m.MaxValue = 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if(m.MaxValue < m.Values[i][j]){
				m.MaxValue = m.Values[i][j]
			}
		}		
	}
}
	
func (m Matrix) print() {
	for i := 0; i < len(m.Values); i++ {
		for j := 0; j < len(m.Values[i]); j++ {
			fmt.Print("|", m.Values[i][j], "|")
		}
		fmt.Println("")
	}
	fmt.Println("Alto: ", m.DimAlto)
	fmt.Println("Ancho: ", m.DimAncho)
	if m.Cuadratica {
		fmt.Println("Es cuadratica")
	} else {
		fmt.Println("No es cuadratica")
	}
	fmt.Println("Valor maximo: ", m.MaxValue)
}

