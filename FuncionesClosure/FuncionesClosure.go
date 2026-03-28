package main

import (
	"fmt"
)

func configuracionDeCobra(cobroPorDia int) func(int) int {
	return func(horas int) int {
		return cobroPorDia * horas
	}
}
func main() {
	cobrarVip := configuracionDeCobra(500)
	fmt.Println("La zona vip se cobra: ", cobrarVip(2))
}
