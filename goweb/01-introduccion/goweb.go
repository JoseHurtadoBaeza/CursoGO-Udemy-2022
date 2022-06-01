package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Crear servidor
	//http.ListenAndServe("localhost:3000", nil)

	fmt.Println("El servidor está corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000/")

	// Con log.Fatal si se devuelve algún error vamos a poder logear dicho error
	log.Fatal(http.ListenAndServe("localhost:3000", nil))

}
