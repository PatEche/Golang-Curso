package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
			var suma int = 0

			for i := 0; i <= 10; i++ {
				if i%2 != 0 {
					fmt.Print(i, ",")
				}

			}


			fmt.Println("La suma de los primeros 10 numeros es:", suma)


		miMapa := map[string]string{
			"Colombia":  "Bogota",
			"Venezuela": "Caracas",
			"Brasil":    "Brasilia",
			"Chile":     "Santiago",
		}

		//For Each in Go
		for k, v := range miMapa {
			fmt.Println("La capital de " + k + " es: " + v)
		}

	*/

	/* Go no tiene ni While ni Do While, lo implementa dentro del for con if al comienzo o al final del for,
	siempre dentro del cuerpo del for. Si evaluamos al inicio equivale a un repita mientras,
	si ejecuta una vez y al final se evalua la condicion equivale a un repita hasta que*/

	var fruta string

	for {

		fmt.Println("Indique una fruta de 4 letras y adivine:")
		fmt.Scan(&fruta)
		fruta = strings.ToLower(fruta)

		if fruta == "pera" {
			fmt.Println("Muy bien la fruta que cumple con la condicion es:", fruta, "y usted ingresó", fruta)
			break
		} else {
			fmt.Println("No es la fruta esperada!")
		}

	}

	// Falta implementar el repita hasta.
}
