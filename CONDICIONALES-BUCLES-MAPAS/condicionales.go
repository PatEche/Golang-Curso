package main

import "fmt"

//Condicionales

func condicionales() {

	var edad int = 40
	var permiso bool = false

	if edad < 18 && permiso {
		fmt.Println("El menor de edad puede viajar solo porque tiene permiso")
	} else if edad <= 18 && !permiso {
		println("El menor de edad no pude viajar solo")
	} else {
		fmt.Println("La persona puede viajar sola porque es mayor de edad")
	}

}
