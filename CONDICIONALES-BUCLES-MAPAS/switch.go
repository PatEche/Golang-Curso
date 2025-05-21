package main

import "fmt"

func condiciones() {

	var fruta string

	fmt.Println("Ingrese un valor alfanumerico:")
	fmt.Scan(&fruta)

	switch fruta {
	case "manzana":
		fmt.Println("Has ingresado manzana.")
	case "pera":
		fmt.Println("Haz ingresado pera.")
	default:
		fmt.Println("No reconozco ese valor.")
	}

}
