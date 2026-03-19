package main

import (
	"fmt"
)

func main() {
	var ganancias [7]float64 // tiene 7 numeros

	var sumatoria float64 = 0
	for a := 0; a < len(ganancias); a++ {
		fmt.Printf("Dia %d: ", a+1) // 1 ,2 ,3
		fmt.Scanln(&ganancias[a])

		sumatoria += ganancias[a]
	}
	promedio := sumatoria / float64(len(ganancias))

	fmt.Println("TOTAL DE GANANCIAS: ", sumatoria)
	fmt.Println("PROMEDIO EN GANANCIA: ", promedio)

	fmt.Println("--GANANCIAS MENOR DE 500--")

	for i := 0; i < len(ganancias); i++ {
		if ganancias[i] < 500 {
			fmt.Printf("El dia %d tuvo un dia bajo con %.2f\n", i+1, ganancias[1])
		}
	}

}
