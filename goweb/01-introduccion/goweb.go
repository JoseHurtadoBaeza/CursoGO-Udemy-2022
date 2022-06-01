package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handlers
func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("El método es + " + r.Method)
	fmt.Fprintln(rw, "Hola Mundo")
}

func PaginaNF(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

func main() {

	// Router
	http.HandleFunc("/", Hola)
	http.HandleFunc("/page", PaginaNF)

	// Crear servidor
	//http.ListenAndServe("localhost:3000", nil)

	fmt.Println("El servidor está corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000/")

	// Con log.Fatal si se devuelve algún error vamos a poder logear dicho error
	log.Fatal(http.ListenAndServe("localhost:3000", nil))

}
