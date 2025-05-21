package main

import "fmt"

func main() {

	// var nombrePersona string = "Patricio"
	// var apellidoPersona = "Echeverria"

	// segundoNombre := "Luis"

	// fmt.Println("Hola Mundo. Soy " + nombrePersona)

	// fmt.Println("Mi apellido es " + apellidoPersona)

	// fmt.Println("Mi segundo nombre es " + segundoNombre)

	/* Parte numerica */

	// var anioActual int16 = 2025

	// edad := 41

	// var edadNew int8 = 90

	// fmt.Println("El año actual es", anioActual)

	// fmt.Println("La edad es", edad)

	// fmt.Println("La edad nueva es", edadNew)

	/* Arreglos y slices */

	// var listaFrutas = [4]string{"Pera", "Naranja", "Manzana", "Banana"}
	// fmt.Println((listaFrutas))

	listaPaises := []string{"Argentina", "Brasil", "España"}
	fmt.Println("Lista original:", listaPaises)

	listaPaises = append(listaPaises, "Chile")
	fmt.Println("Lista con append, agregando un elemento al final del slice:", listaPaises)

	// fmt.Println(listaPaises[1:3]) // Con los ":" obtenemos rangos de elementos del slice

	listaPaises2 := listaPaises[1:]
	fmt.Println("Lista indicando que muestre desde el indice 1 hasta el final:", listaPaises2)

}
