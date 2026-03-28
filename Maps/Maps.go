package main

import (
	"fmt"
)

func main() {
	//Declarar un map
	a := make(map[string]string)
	//Declarar un map con capacidad
	b := make(map[string]string, 7)
	//Darles valor a los maps
	a["Nombre"] = "Ezequiel"
	b["Edad"] = "20"

	nombres := map[string]string{
		"Madre":     "Silvia",
		"Padre":     "Carlos",
		"Hijo":      "Eze",
		"NoviaHijo": "Iara",
	}

	//Eliminar elemento
	delete(nombres, "Madre")
	fmt.Println(nombres)

	//Fijarse si un valor del map no existe
	edades := map[string]int{
		"Eze":  23,
		"Iara": 2,
	}
	if edades, ok := edades["Hola"]; ok {
		if edades > 18 {
			fmt.Println("Es mayor")
		} else {
			fmt.Println("Es menor de edad")
		}
	} else {
		fmt.Println("No existe el valor")
	}

	//Recorrer un map
	for l, edades := range edades {
		fmt.Printf("La edad  de %s es : %d \n ", l, edades)
	}

}
