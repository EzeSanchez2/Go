package main

import (
	"fmt"
)

func CargarSaldo(saldo *int, monto int) {
	*saldo += monto
	fmt.Println("Cargando $... , operacion exitosa")
}

func PagarViaje(saldo *int) {
	*saldo = *saldo - 60
	fmt.Println("Viaje pagado. Saldo descontado de la memoria")
}

func main() {
	miSaldo := 100
	fmt.Println("Saldo inicial: ", miSaldo)
	CargarSaldo(&miSaldo, 500)
	fmt.Println("Tus saldo actual es ", miSaldo)
	PagarViaje(&miSaldo)
	fmt.Println("Saldo final en targeta ", miSaldo)
}
