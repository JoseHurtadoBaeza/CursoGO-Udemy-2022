package main

import (
	"fmt"
	"strings"
)

func sumar(expresion string) {

	// 4+4+5+8
	valores := strings.Split(expresion, "+")

	fmt.Printf("%T \n", valores)

}

func main() {

	var expresion string
	//var result int

	fmt.Printf("=>")
	fmt.Scanln(&expresion)

	sumar(expresion)

}
