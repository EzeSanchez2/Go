package main

import (
	"fmt"
	"strconv"
)

func main() {
	var cadena string
	cadena = "Hola mundo"
	fmt.Println(len(cadena) - 1)

	edad := 25
	fmt.Println("Mi edad es " + strconv.Itoa(edad)) // Lo convierte en string al 25
}
