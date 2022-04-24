package main

import "fmt"

// Closures

func main() {

	// Función anónima
	/*func() {
		fmt.Println("Hola desde la función anónima")
	}()*/

	myfunc := func(nombre string) string {
		return fmt.Sprintf("Hola, %s desde la función anónima", nombre)
	}

	fmt.Println(myfunc("Jose"))

}
