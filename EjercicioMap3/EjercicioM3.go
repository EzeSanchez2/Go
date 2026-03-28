package main

import (
	"fmt"
)

func main() {
	estacionamiento := make(map[string][]string)
	var patente string
	var sector string
	for {
		fmt.Println("INGRESE UNA PATENTE")
		fmt.Scanln(&patente)
		if patente == "SALIR" {
			break
		}
		fmt.Println("INGRESE UN SECTOR : A, B, o C")
		fmt.Scanln(&sector)

		estacionamiento[sector] = append(estacionamiento[sector], patente)
	}

	for sector, listaPatente := range estacionamiento {
		fmt.Println("SECTOR: ", sector)
		for _, p := range listaPatente {
			fmt.Println("MOTO: ", p)
		}
	}

	var sectorVaciar string
	fmt.Println("INGRESE UN SECTOR PARA VACIAR")
	fmt.Scanln(&sectorVaciar)

	delete(estacionamiento, sectorVaciar)

	fmt.Println("SECTOR ELIMINADO. ACTUALMENTE ESTA ASI: ", estacionamiento)

}
