package main

import (
	"fmt"
)

func main() {
	var CantidadNotas int
	fmt.Println("Ingrese la cantidad de notas: ")
	fmt.Scanln(&CantidadNotas)
	divisor := float32(CantidadNotas)
	var sumatoria float32 = 0
	for CantidadNotas > 0 {
		fmt.Println("Ingrese nota del alumno: ")
		var notas float32
		fmt.Scanln(&notas)

		sumatoria += notas

		CantidadNotas--
	}

	promedio := sumatoria / divisor
	fmt.Println("El promedio del alumno es: ", promedio)

}
