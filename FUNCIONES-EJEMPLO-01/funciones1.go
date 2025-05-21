package main

import "fmt"

// FUNCIONES EJEMPLO N1

func suma(a int, b int) int {
	return a + b

}

func resta(a int, b int) int {
	return a - b

}

func imprimeResultadoSumar(a int, b int) {
	fmt.Println("Para a =", a, "y b =", b, "la suma resultante es:", suma(a, b))
}

func imprimeResultadoResta(a int, b int) {
	fmt.Println("Para a =", a, "y b =", b, "la resta resultante es:", resta(a, b))
}

func main() {
	imprimeResultadoSumar(5, 8)
	imprimeResultadoSumar(22, 4)
	imprimeResultadoResta(15, 8)
	imprimeResultadoResta(55, 34)

}
