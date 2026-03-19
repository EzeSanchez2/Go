package main

import (
	"fmt"
)

func main() {
	fmt.Println("--Ingrese 10 temperaturas--")
	var mediciones [10]float64
	for a := 0; a < len(mediciones); a++ {
		fmt.Printf("Temperatura %d: ", a+1)
		fmt.Scanln(&mediciones[a])
	}

	max := mediciones[0]
	min := mediciones[0]
	for b := 0; b < len(mediciones); b++ {

		if mediciones[b] > max {
			max = mediciones[b]
		}
		if mediciones[b] < min {
			min = mediciones[b]
		}
	}
	fmt.Println("--------------------------------------------------")
	fmt.Println("--Estado de tempratura--")
	fmt.Println("  ")
	for i := 0; i < len(mediciones); i++ {
		switch {
		case mediciones[i] < 15:
			fmt.Println("HACE MUCHO FRIO")
		case mediciones[i] >= 15 && mediciones[i] <= 30:
			fmt.Println("TEMPERATURA NORMAL")
		case mediciones[i] > 30:
			fmt.Println("HACE MUCHO CALOR")
		}
	}
	fmt.Println("--------------------------------------------------")

	fmt.Println("LA TEMPERATURA MAXIMA FUE DE: ", max)
	fmt.Println("LA TEMPERATURA MINIMA FUE DE: ", min)

}
