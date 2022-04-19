package main

import (
	"fmt"
	"strings"
)

// Función que devuelve la cadena pasada por parámetro al revés
func reverse(cadena string) string {

	arrayCadena := strings.Split(cadena, " ")
	fmt.Println(arrayCadena)

	return ""
}

func esPalindromo(palabra string) {

	/*fmt.Println(palabra)
	palabra = strings.ToLower(palabra)*/

	fmt.Println(palabra)
	palabra = strings.ToUpper(palabra)
	fmt.Println(palabra)

	//palabra = strings.Replace(palabra, "Z", "S", 2)ç
	palabra = strings.ReplaceAll(palabra, " ", "")
	fmt.Println(palabra)

}

func main() {

	// Una palabra palíndroma es que se lee de igual forma de izquierda a derecha y viceversa
	// Un ejemplo podría ser la palabra "oso"
	//esPalindromo("Oso")
	esPalindromo("Luz azul") // También es una palabra palíndroma si le quitamos ese espacio

	reverse("Luz azul")

}
