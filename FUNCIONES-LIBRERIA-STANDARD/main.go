package main

import (
	"fmt"
)

// "os"

// CONCEPTOS AVANZADOS DE FUNCIONES

// func imprimeResultadosumar(operacion func(int, int) int,
// 	a int, b int) {

// 	fmt.Println("Para a =", a, "y b =", b, "la sumar resultante es:", operacion(a, b))
// }

// func imprimeResultadorestar(operacion func(int, int) int,
// 	a int, b int) {
// 	fmt.Println("Para a =", a, "y b =", b, "la restar resultante es:", operacion(a, b))
// }

// func imprimeResultadoMultiplicacion(operacion func(a int, b int) int,
// 	a int, b int) {
// 	fmt.Println("Para a =", a, "y b =", b, "la multiplicacion resultante es:", operacion(a, b))
// }

var funciones = map[string]func(int, int) int{
	"suma":           func(a int, b int) int { return a + b },
	"resta":          func(a int, b int) int { return a - b },
	"multiplicacion": func(a int, b int) int { return a * b },
	"division":       func(a int, b int) int { return a / b },
}

func imprimeResultado(operacion string, a int, b int) {

	f, exists := funciones[operacion] // En caso de que la operacion no exista manejo esa excepcion!!

	if !exists {
		fmt.Println("Operacion no valida")
		return
	}

	fmt.Println("Para a =", a, "y b =", b, "la", operacion, "resultante es:", f(a, b))
}

func main() {

	imprimeResultado("suma", 10, 6)
	imprimeResultado("resta", 8, 4)
	imprimeResultado("multiplicacion", 8, 4)
	imprimeResultado("division", 8, 4)

}
