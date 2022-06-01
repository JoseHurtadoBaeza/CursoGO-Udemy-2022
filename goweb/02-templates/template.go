package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(rw, "Hola Mundo")
	template, err := template.ParseFiles("index.html") // ParseFiles devuelve dos valores: el template en sí y también un error

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, nil)
	}

}

func main() {

	// Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	// Servidor
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	fmt.Println("El servidor está corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000/")
	log.Fatal(server.ListenAndServe())

}
