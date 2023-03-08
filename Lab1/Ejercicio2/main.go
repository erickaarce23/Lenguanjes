package main

import (
	"fmt"
	"math"
)

func imprimirFigura(numero int) {
	if numero%2 == 0 {
		fmt.Println("El número ingresado no es impar positivo")
		return
	}

	maxEspacios := int(math.Floor(float64(numero) / 2.0))

	for i := 1; i <= numero; i += 2 {
		espacios := maxEspacios - int(math.Floor(float64(i)/2.0))
		estrellas := i

		for j := 0; j < espacios; j++ {
			fmt.Print(" ")
		}

		for j := 0; j < estrellas; j++ {
			fmt.Print("*")
		}

		fmt.Println("")
	}

	for i := numero - 2; i >= 1; i -= 2 {
		espacios := maxEspacios - int(math.Floor(float64(i)/2.0))
		estrellas := i

		for j := 0; j < espacios; j++ {
			fmt.Print(" ")
		}

		for j := 0; j < estrellas; j++ {
			fmt.Print("*")
		}

		fmt.Println("")
	}
}

func main() {
	imprimirFigura(5)
}
