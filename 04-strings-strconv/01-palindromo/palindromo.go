package main

import (
	"fmt"
	"strings"
)

func esPalindromo(palabra string) {

	fmt.Println(palabra)
	palabra = strings.ToLower(palabra)

	fmt.Println(palabra)
	palabra = strings.ToUpper(palabra)
	fmt.Println(palabra)

}

func main() {

	// Una palabra palíndroma es que se lee de igual forma de izquierda a derecha y viceversa
	// Un ejemplo podría ser la palabra "oso"
	esPalindromo("Oso")

}
