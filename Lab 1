/*
Haga un programa que cuente e indique el número de caracteres,
número de palabras y el número de líneas de un texto cualquiera
obtenido de cualquier forma que considere). Considere hacer el
de las tres variables en el mismo proceso."
*/
package main

import (
	"fmt"
	"strings"
)

func main() {

	//Se declara el texto
	text := "Haga un programa, que cuente indique el \n número de caracteres el número de palabras \n y el número de líneas de un texto cualquiera"

	//devuelve un slice de strings separados por " "
	palabras := strings.Split(text, " ")

	//devuelve un slice de strings separados por \n
	lineas := strings.Split(text, "\n")

	count := 0
	//ciclo que recorre el slice palabras y las cuenta
	for i := 0; i < len(palabras); i++ {
		count++
	}

	//recorre el slice lineas devoliendo la cantidad
	conta := len(lineas)

	fmt.Println("El numero de caracteres es: ", len(text))
	fmt.Println("El numero de palabras es: ", count)
	fmt.Println("El numero de lineas es: ", conta)
}
