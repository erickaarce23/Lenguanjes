package main

import "fmt"

func rotate(arr *[]int, n int, dir int) {
	l := len(*arr)
	temp := make([]int, l)

	// En caso de dirección derecha, calcula nueva posición
	if dir == 1 {
		n = l - n
	}

	// Copia los elementos a una matriz temporal
	for i := 0; i < l; i++ {
		temp[i] = (*arr)[i]
	}

	// Gira los elementos a la izquierda
	for i := 0; i < l; i++ {
		j := (i + n) % l
		(*arr)[i] = temp[j]
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	rotate(&arr, 2, 0) // Gira a la izquierda en 2 posiciones

	fmt.Println(arr) // Resultado esperado: [3 4 5 1 2]
}
