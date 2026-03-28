package main

import (
	"fmt"
)

func main() {
	tarifas := map[string]int{
		"Scooter":   20,
		"Deportiva": 50,
		"Custom":    100,
	}

	fmt.Println(tarifas)

	estacionamiento := make(map[string]int)
	var patente string
	var hora int
	for {
		fmt.Println("INGRESE UNA PATENTE")
		fmt.Scanln(&patente)
		fmt.Println("Patente: ", patente)
		if patente == "SALIR" {
			break
		}
		fmt.Println("INGRESE LA HORA")
		fmt.Scanln(&hora)
		fmt.Println("Hora: ", hora)

		estacionamiento[patente] = hora
	}

	for patente, hora := range estacionamiento {
		precioPorHora := tarifas["Deportiva"]
		total := precioPorHora * hora

		fmt.Printf("La patente %q estuvo %d y pago %d  ", patente, hora, total)

	}

	var consulta string
	fmt.Println("INGRESE PATENTE A CONSULTAR")
	fmt.Scanln(&consulta)
	if hora, ok := estacionamiento[consulta]; ok {
		precioPorhora := tarifas["Deportiva"]
		total := precioPorhora * hora
		fmt.Printf("La moto %q debe un total %d\n", consulta, total)
	} else {
		fmt.Println("No se encontro moto en el sistema")
	}

}
