package main

import (
	"fmt"
)

type Vehiculo struct {
	Patente string
	Marca   string
}

type Moto struct {
	Vehiculo
	Cilindrada int
}

type Auto struct {
	Puertas int
}

func main() {
	moto := Moto{
		Vehiculo{
			Patente: "avf324",
			Marca:   "Fiat",
		},
		16,
	}

	auto := Auto{
		5,
	}

	fmt.Println("Moto patente: ", moto.Patente)
	fmt.Println("auto puertas: ", auto.Puertas)
}
