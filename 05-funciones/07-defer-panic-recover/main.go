package main

import (
	"fmt"
	"os"
)

func main() {

	if file, error := os.Open("hola.txt"); error != nil {
		fmt.Println("No fue posible de esta forma")
	} else {
		defer func() { // Se ejecutará al final
			fmt.Println("Cerrando el archivo!")
			file.Close()
		}()
		contenido := make([]byte, 254)
		long, _ := file.Read(contenido) // Devuelve la longitud y el error
		contenidoArchivo := string(contenido[:long])
		fmt.Println(contenidoArchivo)
	}

}
