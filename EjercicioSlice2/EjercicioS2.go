package main

import (
	"fmt"
)

func main() {
	var repuestos int
	stock := make([]int, 0, 10)
	fmt.Println("INGRESE LOS REPUESTOS: ")
	for a := 0; a < 8; a++ {
		fmt.Scanln(&repuestos)
		stock = append(stock, repuestos)
		fmt.Println("Repuesto numero: ", repuestos)

	}

	fmt.Println("**ELIMINAR REPUESTOS**")

	for e := len(stock) - 1; e >= 0; e-- {
		if stock[e] == 0 {
			stock = append(stock[:e], stock[e+1:]...)
		}

	}

	var nuevo = []int{50, 80, 100}
	stock = append(stock, nuevo...)

	fmt.Println("Stock resultante: ", stock)
	fmt.Printf("Capacidad: %d", cap(stock))

}
