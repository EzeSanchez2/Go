package main

import (
	"fmt"
)

func main() {
	nombre := []string{
		"Ezequiel",
		"Kiara",
		"Silvia",
		"Carlos",
		"Iara",
	}

	for index, nombre := range nombre {
		fmt.Printf("El nombre %q esta en el indice %d \n ", nombre, index)
	}

	for index := range nombre {
		fmt.Println(index)
	}

	for _, nombre := range nombre {
		fmt.Println(nombre)
	}
}
