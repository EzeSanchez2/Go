package main

import (
	"fmt"
)

func main() {
	var patente string
	estacionamiento := make([]string, 0, 5)
	for {
		fmt.Println(" **INGRESE PATENTE  O FIN PARA SALIR**")
		fmt.Scanln(&patente)
		if patente != "FIN" {
			estacionamiento = append(estacionamiento, patente)
			fmt.Println("Patente: ", patente)
		} else {
			break
		}

	}
	respaldo := make([]string, len(estacionamiento))
	copy(respaldo, estacionamiento)
	fmt.Println("RESPALDO: \n ", respaldo)

	fmt.Println("**** ELIMINAR PRIMER ELEMENTO****")
	estacionamiento = estacionamiento[1:]
	fmt.Println(estacionamiento)

}
