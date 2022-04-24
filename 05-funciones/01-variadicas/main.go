package main

import "fmt"

// Esto lo hacemos para indicar que no sabemos cuántos valores vamos a recibir en la función
func sumar(numeros ...int) int {

	var total int
	for _, num := range numeros {
		total += num
	}

	return total

}

func main() {

	result := sumar(10, 20, 40, 70, 60)

	fmt.Println(result)

}
