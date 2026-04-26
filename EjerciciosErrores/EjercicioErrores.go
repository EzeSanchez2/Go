package main

import (
	"errors"
	"fmt"
)

var (
	ErrSaldoInsuficuente = errors.New("no tenés suficiente plata en la cuenta")
	ErrMontoNegativo     = errors.New("no podés retirar un monto menor o igual a cero")
	ErrLimiteExcedido    = errors.New("el máximo por retiro es $50.000")
)

func Retirar(saldoActual float64, monto float64) (float64, error) {
	if monto <= 0 {
		return saldoActual, ErrMontoNegativo
	}

	if monto > saldoActual {
		return saldoActual, ErrSaldoInsuficuente
	}

	if monto > 50000 {
		return saldoActual, ErrLimiteExcedido
	}

	NuevoSaldo := saldoActual - monto
	return NuevoSaldo, nil
}

func chequearError(nuevoSaldo float64, err error) {
	if err != nil {
		fmt.Println("ERROR BANCARIO: ", err)
	} else {
		fmt.Println("RETIRO EXITOSO. NUEVO SALDO: ", nuevoSaldo)
	}
}

func main() {
	saldo, err := Retirar(100000, 20000)
	chequearError(saldo, err)

	saldoA, errA := Retirar(1000, 89080800)
	chequearError(saldoA, errA)

}
