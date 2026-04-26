package main

import (
	"errors"
	"fmt"
)

type Vehiculo struct {
	patente string
	multas  int
}

type Moto struct {
	Vehiculo
	cilindtadas int
}

var (
	ErrAccesoDenegado = errors.New("El vehículo tiene multas pendientes y no puede entrar")
)

func (v *Vehiculo) AgregarMultas() {
	v.multas = v.multas + 1
}

type Controlador interface {
	ValidarAcceso() error
}

func (m Moto) ValidarAcceso() error {
	if m.multas > 0 {
		return ErrAccesoDenegado
	}
	return nil
}

func CheckearEntrada(c Controlador) {
	err := c.ValidarAcceso()
	if err != nil {
		fmt.Println("ALERTA!, ", err)
	} else {
		fmt.Println("TODO ESTA BIEN!")
	}
}

func main() {
	miMoto := Moto{
		Vehiculo: Vehiculo{
			patente: "ABC7890",
			multas:  0,
		},
		cilindtadas: 120,
	}

	CheckearEntrada(miMoto)
}
