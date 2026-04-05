package main

import (
	"fmt"
)

type Cuenta struct {
	Titular string
	Saldo   float64
}

func (c Cuenta) SimularPrestamo(cuotas int) float64 {
	prestamoSimulado := c.Saldo / float64(cuotas)
	fmt.Println("El prestamo simulado es: ", prestamoSimulado)
	return prestamoSimulado
}

func (r *Cuenta) Depositar(monto float64) {
	r.Saldo += monto
}

func main() {
	cuenta := Cuenta{
		Titular: "Ezequiel Sanchez",
		Saldo:   10000,
	}
	cuenta.SimularPrestamo(10)
	cuenta.Depositar(5000)

	fmt.Println("El saldo es: ", cuenta.Saldo)

}
