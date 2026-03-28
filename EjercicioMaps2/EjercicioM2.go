package main

import (
	"fmt"
)

func main() {
	conteoMarcas := make(map[string]int)
	var marca string
	for {
		fmt.Println("INGRESE UNA MARCA")
		fmt.Scanln(&marca)

		if marca == "SALIR" {
			break
		}
		conteoMarcas[marca]++
	}

	for marca, cantidad := range conteoMarcas {
		fmt.Printf("Hay %d cantidad de la marca %q\n", cantidad, marca)
	}

	fmt.Println(conteoMarcas)

}
