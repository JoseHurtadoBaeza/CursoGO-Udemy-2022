package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expresion string) int {

	// 4+4+5+8
	valores := strings.Split(expresion, "+")

	var result int

	for _, valor := range valores {
		num, _ := strconv.Atoi(valor)

		result += num
	}

	return result

}

func main() {

	var expresion string
	var result int

	fmt.Printf("=>")
	fmt.Scanln(&expresion)

	result = sumar(expresion)

	fmt.Printf("=> %d\n", result)

}
