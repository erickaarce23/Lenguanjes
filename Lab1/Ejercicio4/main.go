package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type persona struct {
	nombre string
	edad   int
}

type Personas []persona //slice tipo persona

var listaPersonas Personas

func (f *Personas) agregarPersona(nom string, edad int) { //agrega personas a la lista Personas

	indice := slices.IndexFunc(*f, func(p persona) bool { return p.nombre == nom })
	if indice == -1 {
		*f = append(*f, persona{nom, edad})
	} else {
		fmt.Println("cant add existing person to the list")
	}
}

func map1[P1, P2 any](lista []P1, f func(P1) P2) []P2 {
	mapped := make([]P2, len(lista))
	for i, e := range lista {
		mapped[i] = f(e)
	}
	return mapped
}

func filter1[T1 any](lista []T1, f func(T1) bool) []T1 {
	filtered := make([]T1, 0)

	for _, elemento := range lista {
		if f(elemento) {
			filtered = append(filtered, elemento)
		}
	}
	return filtered
}

func (p *Personas) personasMayoresEdad() []persona {
	return filter1(*p, func(p persona) bool {
		return p.edad >= 65
	})
}

func main() {
	listaPersonas.agregarPersona("Pepe", 20)
	listaPersonas.agregarPersona("Juana", 66)
	listaPersonas.agregarPersona("Pedro", 82)
	listaPersonas.agregarPersona("Lula", 18)

	println(" ")
	fmt.Println(listaPersonas.personasMayoresEdad())
	println(" ")
}
