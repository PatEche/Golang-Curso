package main

import (
	"fmt"
	// "strings"
)

// CONCEPTOS AVANZADOS DE FUNCIONES

func imprimeResultadoSuma(operacion func(int, int) int,
	a int, b int) {

	fmt.Println("Para a =", a, "y b =", b, "la suma resultante es:", operacion(a, b))
}

func imprimeResultadoResta(operacion func(int, int) int,
	a int, b int) {
	fmt.Println("Para a =", a, "y b =", b, "la resta resultante es:", operacion(a, b))
}

func main() {

	suma := func(a int, b int) int { //Funcion anonima o closure
		return a + b

	}

	resta := func(a int, b int) int {
		return a - b

	}

	imprimeResultadoSuma(suma, 10, 5) // Como primer parametro paso la funcion, seguido los valores para calcular
	imprimeResultadoResta(resta, 10, 4)
}
