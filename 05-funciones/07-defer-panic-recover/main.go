package main

import (
	"fmt"
	"os"
)

func main() {

	file, error := os.Open("hola.txt")

	if error != nil {
		fmt.Println("No fue posible de esta forma")
	} else {
		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:long])
		fmt.Println(contenidoArchivo)
	}

	file.Close()

}
