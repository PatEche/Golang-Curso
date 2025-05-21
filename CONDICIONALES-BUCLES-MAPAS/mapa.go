package main

import "fmt"

func mapa() {
	miMapa := map[string]string{
		"Colombia":  "Bogota",
		"Venezuela": "Caracas",
		"Brasil":    "Brasilia",
		"Chile":     "Santiago",
	}

	fmt.Println("El mapa de paises es:", miMapa)

	fmt.Println("La capital de Venezuela es:", miMapa["Venezuela"])

	miMapa["Argentina"] = "Buenos Aires"

	fmt.Println("La capital de Argentina es:", "Argentina")

	delete(miMapa, "Venezuela")

	fmt.Println("El mapa de paises es:", miMapa)

	fmt.Println("La capital de Venezuela es:", miMapa["Venezuela"])

}
