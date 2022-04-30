package main

import "fmt"

// Struct persona
type Persona struct {
	nombre string
	edad   int
}

func main() {

	p1 := Persona{"Jose", 25}

	fmt.Println(p1)

	p1.nombre = "Carlos"

	fmt.Println(p1)

	p2 := Persona{
		nombre: "Pepe",
		edad:   62,
	}

	fmt.Println(p2)

}
