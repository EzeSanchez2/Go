package main

import "fmt"

func main() {
	var numero int // Formal extensa
	println(numero)

	println(numero * 5)

	numero = 25
	fmt.Println(numero)

	numero = 40
	fmt.Println(numero)

	nombre := "Alejandro" // Aca con el := decalramos indirectamente la variable como String
	nombre = "Eze"        // Aca ya lo detecta
	fmt.Println(nombre)   // Imprime "Eze"

	nombre, numero = "Alejandro", 15
	fmt.Println(nombre, numero)

	nombre3, numero3 := "Ezequiel", 23
	fmt.Println(nombre3, numero3)

	var ( // Declarar variables de forma general
		personaje          = "Goku"
		enemigo1, enemigo2 = "Freezer", "Cell"
	)
	fmt.Println(personaje, enemigo1, enemigo2)
}
