package main

import (
	"fmt"
)

type Cobrador interface {
	CalcularComision(monto float64) float64
}

type UsuarioPlata struct {
	nombre string
}

type UsuarioOro struct {
	nombreO   string
	descuento float64
}

func (u UsuarioPlata) CalcularComision(monto float64) float64 {
	return monto * 0.10
}

func (u UsuarioOro) CalcularComision(monto float64) float64 {
	comisionBase := monto * 0.5
	return monto - comisionBase
}

func (u *UsuarioOro) ActualizarDescuentoNuevo(descuentoNuevo float64) {
	u.descuento = descuentoNuevo
}

func ProcesarPagos(c Cobrador, monto float64) {
	resultado := c.CalcularComision(monto)
	fmt.Println("Comision Calculada: ", resultado)
}

func main() {
	usuarioP := UsuarioPlata{
		nombre: "Ezequiel",
	}

	usuarioO := UsuarioOro{
		nombreO:   "Iara",
		descuento: 0,
	}

	usuarioO.ActualizarDescuentoNuevo(100)
	ProcesarPagos(usuarioP, 5000)
	ProcesarPagos(usuarioO, 5000)
}
