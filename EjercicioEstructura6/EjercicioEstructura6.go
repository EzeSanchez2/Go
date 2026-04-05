package main

import (
	"fmt"
)

type Moto struct {
	Patente  string
	EnCasino bool
}

type Estacionamiento struct {
	Nombre    string
	Capacidad int
	Motos     []Moto
}

func (e *Estacionamiento) IngresarMoto(m Moto) {
	e.Motos = append(e.Motos, m)
	m.EnCasino = true
}

func main() {
	estacionamiento := Estacionamiento{
		Nombre:    "Parkin Central Casino",
		Capacidad: 5,
		Motos:     []Moto{},
	}

	m1 := Moto{Patente: "ABC-123", EnCasino: false}
	m2 := Moto{Patente: "DEF-456", EnCasino: false}
	m3 := Moto{Patente: "GHI-789", EnCasino: false}

	estacionamiento.IngresarMoto(m1)
	estacionamiento.IngresarMoto(m2)
	estacionamiento.IngresarMoto(m3)

	fmt.Println("----------------------------------------------------------------")

	fmt.Println("TOTAL DE MOTOS ADENTRO: ", len(estacionamiento.Motos))
}
