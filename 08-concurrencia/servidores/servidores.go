package main

import (
	"fmt"
	"net/http"
	"time"
)

// Esta función se va a encargar de comprobar si un servidor está funcionando o no
// Recibe la url del servidor por parámetro
func revisarServidor(servidor string) {
	_, err := http.Get(servidor)

	if err != nil {
		fmt.Println("No está disponible")
	} else {
		fmt.Println(servidor, "Está funcionando")
	}

}

func main() {

	inicio := time.Now()

	servidores := []string{
		"https://oregoom.com/",
		"https://www.udemy.com/",
		"https://www.youtube.com/",
		"https://www.facebook.com/",
		"https://www.google.com/",
	}

	for _, servidor := range servidores {
		//revisarServidor(servidor) // Ejecución de forma secuencial
		go revisarServidor(servidor) // Ejecución de forma concurrente
		// Con la palabra reservada go indicamos la creación de múltiples hilos que nos van
		// a permitir ejecutar esta función al mismo tiempo.
	}

	tiempoPaso := time.Since(inicio)

	fmt.Println("Tiempo de ejecución: ", tiempoPaso)

}
