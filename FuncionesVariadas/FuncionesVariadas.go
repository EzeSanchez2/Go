package main

import (
	"fmt"
)

func sum(numeros ...int) int {
	resultado := 0
	for _, num := range numeros {
		resultado += num
	}
	return resultado
}

func cadenasTextos(cadena string, cadenas ...string) string {
	for _, c := range cadenas {
		cadena += " "
		cadena += c
	}

	return cadena
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println("Hola", "Soy", "Eze")
}
